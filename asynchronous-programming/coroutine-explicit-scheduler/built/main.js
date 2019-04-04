"use strict";
function* co_timeline1() {
    yield '1a';
    yield '1b';
    yield* batchTimeline(co_timeline2(), co_timeline5());
    yield '1c';
}
function* co_timeline2() {
    yield '2a';
    yield* batchTimeline(co_timeline3(), co_timeline4());
    yield '2b';
    yield '2c';
}
function* co_timeline3() {
    yield '3a';
    yield '3b';
    yield '3c';
}
function* co_timeline4() {
    yield '4a';
    yield '4b';
    yield '4c';
}
function* co_timeline5() {
    yield '5a';
    yield '5c';
}
function scheduler(timeline) {
    while (true) {
        // every step, the control always transfer back to scheduler
        let step = timeline.next();
        if (step.done) {
            return;
        }
        // collected everything generated in this step
        // this is a perfect place to do thing in large batch
        console.log(step.value);
    }
}
function* batchTimeline(...timelines) {
    let activeTimelines = [];
    while (timelines.length > 0) {
        let batch = [];
        for (let timeline of timelines) {
            let step = timeline.next();
            if (!step.done) {
                activeTimelines.push(timeline);
                batch.push(step.value);
            }
        }
        if (batch.length === 0) {
            return;
        }
        yield batch.join(', ');
        timelines = activeTimelines;
        activeTimelines = [];
    }
}
scheduler(co_timeline1());
//# sourceMappingURL=main.js.map