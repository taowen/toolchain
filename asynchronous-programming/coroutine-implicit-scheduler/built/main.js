"use strict";
function pause() {
    return new Promise(resolve => resolve());
}
let batch = [];
let done = false;
async function co_timeline1() {
    batch.push('1a');
    await pause();
    batch.push('1b');
    await pause();
    await Promise.all([co_timeline2(), co_timeline5()]);
    batch.push('1c');
    done = true;
}
async function co_timeline2() {
    batch.push('2a');
    await pause();
    await Promise.all([co_timeline3(), co_timeline4()]);
    batch.push('2b');
    await pause();
    batch.push('2c');
}
async function co_timeline3() {
    batch.push('3a');
    await pause();
    batch.push('3b');
    await pause();
    batch.push('3c');
}
async function co_timeline4() {
    batch.push('4a');
    await pause();
    batch.push('4b');
    await pause();
    batch.push('4c');
}
async function co_timeline5() {
    batch.push('5a');
    await pause();
    batch.push('5b');
}
async function co_print() {
    while (!done) {
        console.log(batch.join(', '));
        batch = [];
        await pause();
    }
}
co_timeline1();
co_print();
//# sourceMappingURL=main.js.map