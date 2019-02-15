define("lib", ["require", "exports"], function (require, exports) {
    "use strict";
    exports.__esModule = true;
    exports.data = [1, 2, 2, 3, 4, 5, 5, 5, 6];
});
define("main", ["require", "exports", "lib", "uniq"], function (require, exports, lib_1, unique) {
    "use strict";
    exports.__esModule = true;
    console.log(unique(lib_1.data));
});
