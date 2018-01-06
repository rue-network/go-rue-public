// Copyright 2015 The go-ruereum Authors
// This file is part of the go-ruereum library.
//
// The go-ruereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ruereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ruereum library. If not, see <http://www.gnu.org/licenses/>.

// package web3ext contains grue specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"admin":      Admin_JS,
	"chequebook": Chequebook_JS,
	"clique":     Clique_JS,
	"debug":      Debug_JS,
	"rue":        Eth_JS,
	"miner":      Miner_JS,
	"net":        Net_JS,
	"personal":   Personal_JS,
	"rpc":        RPC_JS,
	"shh":        Shh_JS,
	"swarmfs":    SWARMFS_JS,
	"txpool":     TxPool_JS,
}

const Chequebook_JS = `
web3._extend({
	property: 'chequebook',
	methods: [
		new web3._extend.method({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.method({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const Clique_JS = `
web3._extend({
	property: 'clique',
	methods: [
		new web3._extend.method({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.method({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.method({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.method({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const Admin_JS = `
web3._extend({
	property: 'admin',
	methods: [
		new web3._extend.method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.method({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.method({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.method({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.method({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const Debug_JS = `
web3._extend({
	property: 'debug',
	methods: [
		new web3._extend.method({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.method({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new web3._extend.method({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.method({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.method({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new web3._extend.method({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.method({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.method({
			name: 'metrics',
			call: 'debug_metrics',
			params: 1
		}),
		new web3._extend.method({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.method({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.method({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.method({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new web3._extend.method({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.method({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.method({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.method({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.method({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.method({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.method({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.method({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.method({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.method({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.method({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.method({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.method({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.method({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.method({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.method({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.method({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.method({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.method({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.method({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.method({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.method({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.method({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
	],
	properties: []
});
`

const Eth_JS = `
web3._extend({
	property: 'rue',
	methods: [
		new web3._extend.method({
			name: 'sign',
			call: 'rue_sign',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.method({
			name: 'resend',
			call: 'rue_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		}),
		new web3._extend.method({
			name: 'signTransaction',
			call: 'rue_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.method({
			name: 'submitTransaction',
			call: 'rue_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.method({
			name: 'getRawTransaction',
			call: 'rue_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.method({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'rue_getRawTransactionByBlockHashAndIndex' : 'rue_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'rue_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(web3._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`

const Miner_JS = `
web3._extend({
	property: 'miner',
	methods: [
		new web3._extend.method({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.method({
			name: 'stop',
			call: 'miner_stop'
		}),
		new web3._extend.method({
			name: 'setRuerbase',
			call: 'miner_setRuerbase',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.method({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new web3._extend.method({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.method({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
	],
	properties: []
});
`

const Net_JS = `
web3._extend({
	property: 'net',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const Personal_JS = `
web3._extend({
	property: 'personal',
	methods: [
		new web3._extend.method({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new web3._extend.method({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.method({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new web3._extend.method({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.method({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RPC_JS = `
web3._extend({
	property: 'rpc',
	methods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const Shh_JS = `
web3._extend({
	property: 'shh',
	methods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SWARMFS_JS = `
web3._extend({
	property: 'swarmfs',
	methods:
	[
		new web3._extend.method({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.method({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.method({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxPool_JS = `
web3._extend({
	property: 'txpool',
	methods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = web3._extend.utils.toDecimal(status.pending);
				status.queued = web3._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`
