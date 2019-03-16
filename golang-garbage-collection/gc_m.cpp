void runtime·gc_m(void) {
    a.start_time = (uint64)(g->m->scalararg[0]) | ((uint64)(g->m->scalararg[1]) << 32);
    a.eagersweep = g->m->scalararg[2];
    gc(&a);
}

static void gc(struct gc_args *args) {
    // 如果前次回收的清理操作未完成，那么先把这事结束了。
    while (runtime·sweepone() != -1)
    runtime·sweep.npausesweep++;
    // 为回收操作准备相关环境状态。
    runtime·mheap.gcspans = runtime·mheap.allspans;
    runtime·work.spans = runtime·mheap.allspans;
    runtime·work.nspan = runtime·mheap.nspan;
    runtime·work.nwait = 0;
    runtime·work.ndone = 0;
    runtime·work.nproc = runtime·gcprocs();

    //初始化并⾏行标记状态对象 markfor。
    //使⽤用 nproc 个线程执⾏行并⾏行标记任务。
    //任务总数 = 固定内存段(RootCount) + 当前 goroutine G 的数量。 标记函数 markroot。
    runtime·parforsetup(runtime·work.markfor, runtime·work.nproc,
            RootCount + runtime·allglen, nil, false, markroot);

    if (runtime·work.nproc > 1) { // 重置结束标记。
        runtime·noteclear(&runtime·work.alldone);
        // 唤醒 nproc - 1 个线程准备执⾏行 markroot 函数，因为当前线程也会参与标记⼯工作。
        runtime·helpgc(runtime·work.nproc);
    }

    // 让当前线程也开始执⾏行标记任务。
    gchelperstart();
    runtime·parfordo(runtime·work.markfor);
    scanblock(nil, 0, nil);
    if (runtime·work.nproc > 1)
    // 休眠，等待标记全部结束。
    runtime·notesleep(&runtime·work.alldone);
    // 收缩 stack 内存。
    runtime·shrinkfinish();


    // 更新所有 cache 统计参数。
    cachestats();

    // 计算上⼀一次回收后 heap_alloc ⼤大⼩小。
    // 当前 next_gc = heap0 + heap0 * (gcpercent/100)
    // 那么 heap0 = next_gc / (1 + gcpercent/100)
    heap0 = mstats.next_gc*100/(runtime·gcpercent+100);

    // 计算下⼀一次 next_gc 阈值。
    // 这个值只是预估，会随着清理操作⽽而改变。
    mstats.next_gc = mstats.heap_alloc + mstats.heap_alloc * runtime·gcpercent / 100;
    runtime·atomicstore64(&mstats.last_gc, runtime·unixnanotime());


    // ⽬目标是 heap.allspans ⾥里的所有 span 对象。
    runtime·mheap.gcspans = runtime·mheap.allspans; // GC 使⽤用递增的代龄来表⽰示 span 当前回收状态。
    runtime·mheap.sweepgen += 2;
    runtime·mheap.sweepdone = false;
    runtime·work.spans = runtime·mheap.allspans;
    runtime·work.nspan = runtime·mheap.nspan;
    runtime·sweep.spanidx = 0;

    // 并发清理
    if (ConcurrentSweep && !args->eagersweep) {
        // 新建或唤醒⽤用于清理操作的 goroutine。
        if (runtime·sweep.g == nil){
            runtime·sweep.g = runtime·newproc1(&bgsweepv, nil, 0, 0, gc);
        }else if (runtime·sweep.parked) {
            runtime·sweep.parked = false;
            runtime·ready(runtime·sweep.g); // 唤醒 }
        } else { // 串⾏行回收
            // ⽴立即执⾏行清理操作。 while(runtime·sweepone() != -1)
            runtime·sweep.npausesweep++;
        }
    }
}