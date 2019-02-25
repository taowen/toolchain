var voteUpBtn = document.querySelector('#vote_up');
var voteDownBtn = document.querySelector('#vote_down');

function show(votes) {
  document.querySelector('#votes').innerHTML = votes;
}

var voteAction = new Rx.Subject(0);
var voteStrm = voteAction
  .startWith(parseInt(document.querySelector('#votes').innerHTML))
  .scan(function (m, n) {
    return Math.max(m + n, 0);
  });

Rx.Observable.fromEvent(voteUpBtn, 'click')
.mapTo(1)
.subscribe(voteAction);

Rx.Observable.fromEvent(voteDownBtn, 'click')
.mapTo(-1)
.subscribe(voteAction);

voteStrm // Team Few
.filter(function (votes) { return votes < 1000; })
.subscribe(show);

voteStrm // Team Many
.filter(function (votes) { return votes >= 1000; })
.map(function(votes) { return (votes / 1000).toFixed(2) + 'k'; })
.subscribe(show);
