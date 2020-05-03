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
/******/ 	return __webpack_require__(__webpack_require__.s = "./src/app.ts");
/******/ })
/************************************************************************/
/******/ ({

/***/ "./src/app.ts":
/*!********************!*\
  !*** ./src/app.ts ***!
  \********************/
/*! no exports provided */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _client__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./client */ \"./src/client.ts\");\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\n\nconst videos = new Map();\nfunction f() {\n    return __awaiter(this, void 0, void 0, function* () {\n        try {\n            const c = new _client__WEBPACK_IMPORTED_MODULE_0__[\"default\"]();\n            yield c.joinRoom('hoge');\n            c.onTrack = ((clientId, stream) => {\n                const elId = `#partner-${clientId}`;\n                const pre = document.getElementById(elId);\n                pre === null || pre === void 0 ? void 0 : pre.parentNode.removeChild(pre);\n                const $video = document.createElement('video');\n                $video.id = elId;\n                const pa = document.querySelector('#partners');\n                pa.appendChild($video);\n                $video.srcObject = stream;\n                $video.volume = 0;\n                $video.play();\n            });\n        }\n        catch (e) {\n            document.querySelector('#error').innerHTML = e;\n        }\n    });\n}\nconst btn = document.querySelector('button');\nbtn.onclick = f;\n\n\n//# sourceURL=webpack:///./src/app.ts?");

/***/ }),

/***/ "./src/client.ts":
/*!***********************!*\
  !*** ./src/client.ts ***!
  \***********************/
/*! exports provided: default */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, \"default\", function() { return Client; });\nvar __awaiter = (undefined && undefined.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nclass Client {\n    constructor(options) {\n        this.pcs = new Map();\n        this.connected = false;\n        this.endpoint = (options === null || options === void 0 ? void 0 : options.SignalingServerEndpoint) || 'ws://localhost:5000';\n    }\n    createNewPeer(clientId) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const pc = new RTCPeerConnection();\n            this.pcs.set(clientId, pc);\n            if (!this.stream) {\n                const stream = yield navigator.mediaDevices.getUserMedia({\n                    video: true,\n                    audio: true,\n                });\n                this.stream = stream;\n                const $video = document.querySelector('#my-screen');\n                $video.srcObject = stream;\n                $video.volume = 0;\n                yield $video.play();\n            }\n            for (const track of this.stream.getTracks()) {\n                pc.addTrack(track, this.stream);\n            }\n            // TODO 複数のVideo対応する\n            if ('ontrack' in pc) {\n                pc.ontrack = (ev) => __awaiter(this, void 0, void 0, function* () {\n                    this.onTrack(clientId, ev.streams[0]);\n                });\n            }\n            else {\n                // @ts-ignore\n                pc.onaddstream = ((ev) => __awaiter(this, void 0, void 0, function* () {\n                    this.onTrack(clientId, ev.stream);\n                }));\n            }\n            return pc;\n        });\n    }\n    joinRoom(roomID) {\n        return __awaiter(this, void 0, void 0, function* () {\n            return new Promise(((resolve, reject) => {\n                this.ws = new WebSocket(`${this.endpoint}/connect`);\n                this.ws.onmessage = (e) => __awaiter(this, void 0, void 0, function* () {\n                    yield this.handleMessage(JSON.parse(e.data));\n                });\n                this.ws.onopen = (e) => {\n                    this.connected = true;\n                    this.ws.send(JSON.stringify({\n                        type: 'join',\n                        payload: {\n                            room_id: roomID\n                        }\n                    }));\n                    resolve();\n                };\n            }));\n        });\n    }\n    handleMessage(data) {\n        return __awaiter(this, void 0, void 0, function* () {\n            switch (data.type) {\n                case MessageType.NotifyClientId:\n                    this.handleMessageNotifyClientId(data.payload);\n                    break;\n                case MessageType.NewClient:\n                    this.handleNewClient(data.payload);\n                    break;\n                case MessageType.Offer:\n                    yield this.handleMessageOffer(data.payload);\n                    break;\n                case MessageType.Answer:\n                    yield this.handleMessageAnswer(data.payload);\n                    break;\n            }\n        });\n    }\n    handleMessageNotifyClientId(payload) {\n        this.clientId = payload.client_id;\n    }\n    handleNewClient(payload) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const clientId = payload.client_id;\n            const pc = yield this.createNewPeer(clientId);\n            yield this.createOffer(clientId);\n            pc.onicecandidate = (ev) => __awaiter(this, void 0, void 0, function* () {\n                if (!ev.candidate) {\n                    this.sendOffer(clientId);\n                }\n            });\n        });\n    }\n    createOffer(clientId) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const pc = this.pcs.get(clientId);\n            const offer = yield pc.createOffer();\n            yield pc.setLocalDescription(offer);\n            // Trickle ICEならここで初期ICEを送る\n        });\n    }\n    sendOffer(clientId) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const pc = this.pcs.get(clientId);\n            this.ws.send(JSON.stringify({\n                type: MessageType.Offer,\n                payload: {\n                    sdp: pc.localDescription.sdp,\n                    client_id: clientId,\n                }\n            }));\n        });\n    }\n    handleMessageOffer(payload) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const clientId = payload.client_id;\n            const pc = yield this.createNewPeer(clientId);\n            yield pc.setRemoteDescription({\n                type: \"offer\",\n                sdp: payload.sdp,\n            });\n            yield this.createAnswer(clientId);\n            pc.onicecandidate = (ev) => __awaiter(this, void 0, void 0, function* () {\n                if (!ev.candidate) {\n                    this.sendAnswer(clientId);\n                }\n            });\n        });\n    }\n    createAnswer(clientId) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const pc = this.pcs.get(clientId);\n            const answer = yield pc.createAnswer();\n            yield pc.setLocalDescription(answer);\n            // Trickle ICEならここで初期ICEを送る\n        });\n    }\n    sendAnswer(clientId) {\n        const pc = this.pcs.get(clientId);\n        this.ws.send(JSON.stringify({\n            type: MessageType.Answer,\n            payload: {\n                sdp: pc.localDescription.sdp,\n                client_id: clientId,\n            }\n        }));\n    }\n    handleMessageAnswer(payload) {\n        return __awaiter(this, void 0, void 0, function* () {\n            const pc = this.pcs.get(payload.client_id);\n            yield pc.setRemoteDescription({\n                type: \"answer\",\n                sdp: payload.sdp,\n            });\n        });\n    }\n}\nvar MessageType;\n(function (MessageType) {\n    MessageType[\"NotifyClientId\"] = \"notify-client-id\";\n    MessageType[\"NewClient\"] = \"new-client\";\n    MessageType[\"Error\"] = \"error\";\n    MessageType[\"Offer\"] = \"offer\";\n    MessageType[\"Answer\"] = \"answer\";\n})(MessageType || (MessageType = {}));\n\n\n//# sourceURL=webpack:///./src/client.ts?");

/***/ })

/******/ });