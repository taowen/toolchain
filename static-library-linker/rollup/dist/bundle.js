(function () {
  'use strict';

  function unique_pred(list, compare) {
    var ptr = 1
      , len = list.length
      , a=list[0], b=list[0];
    for(var i=1; i<len; ++i) {
      b = a;
      a = list[i];
      if(compare(a, b)) {
        if(i === ptr) {
          ptr++;
          continue
        }
        list[ptr++] = a;
      }
    }
    list.length = ptr;
    return list
  }

  function unique_eq(list) {
    var ptr = 1
      , len = list.length
      , a=list[0], b = list[0];
    for(var i=1; i<len; ++i, b=a) {
      b = a;
      a = list[i];
      if(a !== b) {
        if(i === ptr) {
          ptr++;
          continue
        }
        list[ptr++] = a;
      }
    }
    list.length = ptr;
    return list
  }

  function unique(list, compare, sorted) {
    if(list.length === 0) {
      return list
    }
    if(compare) {
      if(!sorted) {
        list.sort(compare);
      }
      return unique_pred(list, compare)
    }
    if(!sorted) {
      list.sort();
    }
    return unique_eq(list)
  }

  var uniq = unique;

  // 引用的 uniq 包会被链接到最终的 executable 里
  var data = [1, 2, 2, 3, 4, 5, 5, 5, 6];
  console.log(uniq(data));

}());
//# sourceMappingURL=bundle.js.map
