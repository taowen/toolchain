/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/main.js");
/******/ })
/************************************************************************/
/******/ ({

/***/ "../../node_modules/uniq/uniq.js":
/*!******************************************************************!*\
  !*** /home/xiaoju/workspace/toolchain/node_modules/uniq/uniq.js ***!
  \******************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
eval("\n\nfunction unique_pred(list, compare) {\n  var ptr = 1\n    , len = list.length\n    , a=list[0], b=list[0]\n  for(var i=1; i<len; ++i) {\n    b = a\n    a = list[i]\n    if(compare(a, b)) {\n      if(i === ptr) {\n        ptr++\n        continue\n      }\n      list[ptr++] = a\n    }\n  }\n  list.length = ptr\n  return list\n}\n\nfunction unique_eq(list) {\n  var ptr = 1\n    , len = list.length\n    , a=list[0], b = list[0]\n  for(var i=1; i<len; ++i, b=a) {\n    b = a\n    a = list[i]\n    if(a !== b) {\n      if(i === ptr) {\n        ptr++\n        continue\n      }\n      list[ptr++] = a\n    }\n  }\n  list.length = ptr\n  return list\n}\n\nfunction unique(list, compare, sorted) {\n  if(list.length === 0) {\n    return list\n  }\n  if(compare) {\n    if(!sorted) {\n      list.sort(compare)\n    }\n    return unique_pred(list, compare)\n  }\n  if(!sorted) {\n    list.sort()\n  }\n  return unique_eq(list)\n}\n\nmodule.exports = unique\n\n\n//# sourceURL=webpack:////home/xiaoju/workspace/toolchain/node_modules/uniq/uniq.js?");

/***/ }),

/***/ "./src/main.js":
/*!*********************!*\
  !*** ./src/main.js ***!
  \*********************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

eval("// 引用的 uniq 包会被链接到最终的 executable 里\nvar unique = __webpack_require__(/*! uniq */ \"../../node_modules/uniq/uniq.js\");\nvar data = [1, 2, 2, 3, 4, 5, 5, 5, 6];\nconsole.log(unique(data));\n\n//# sourceURL=webpack:///./src/main.js?");

/***/ })

/******/ });