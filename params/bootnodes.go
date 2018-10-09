// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{

	"enode://c723f983d0b9e44d372ac0f9ea33428116b9d9d687c63e01e0cf073b1a4136811b2360ce60ffddfd4acdd9b9f289a53fd0e4f6e8b1b8add337e182e9a437f00b@45.76.206.141:62688",
	"enode://ec0edb0a28fb259cffcb8079d9375efaa0259eaf3c28c0013e00cdb4793e59de2a3b8c61cc76a15f231cd4bea813380d4e7450aa6deeb5f9f7fab62dc7c590ae@217.182.200.54:62688",
	"enode://4f3099a8be7e458c62086e807810e8dd82115488f6b435e7b6cc705a0be936ee3c872bdced1bb30f40ab93f4c84ca4a05e773b99fba6390c8574da1c7864e207@145.239.71.79:62688",
	"enode://e0e6a74159d26bb585ea6069ed71f651b6beffafa2f126b135f79ac00ec029a0508a2f25a615edbde9bcd806c72dc9acec6f1340d92775b746de012c97077d40@195.181.213.118:62688",
	"enode://8b493ec695c58a40ba924795ee0e65b1d24f0a1c9284f3b599be06c95a1c4c7d8128f935748d317d163eac4a5d30c6ea3bc49f1bce6bca64b9bd52a8211c8706@85.255.9.89:62688",
	"enode://2b0ded0e94ea7fdbe4f665403e765262766fe0a692fbe3edca364b35ea851d9edcb8f3c513514cf83eee1ac8753001abc7a2cdc87387255124c52f3df0188047@81.2.252.58:62688",
	"enode://dfe05b2caffb03b6ec155c28e814d65e5155611ef4c2cf19ebea5736419bd3fe20a4a33383a28f02d12831947df4cb9d9ff770f830148dfda50f98c093b7d7d4@85.255.3.94:62688",
	"enode://c411e328ce18848efb036d00b0a3f58d48a60379d068c2984c5808a6a4d0445d4531ee9feb09848394aec4179881902353d38309196c487e90cd0d8afe47571b@54.36.172.209:62688",
	"enode://d2fc2dade2689692f0cb47a6bcd41ae5138cf32cc80e98bb67525794f30dc6843a9bc6678652dbd9c55cc5cff5fea3feef6502f4dd2303ac5cc9015b2e2e35ef@45.76.241.43:62688",
	"enode://bbbb6ed932a7af0dc2289fbc4985428dfb184884599a90143d02a251a7da5f3c4c2be61db3034ad20ae92b7c42b2d2434133edc0a174cc08dbb0bffa1acee1fb@108.61.192.135:62688",
	
	/*// Ethereum Foundation Go Bootnodes

	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Ethereum Foundation Cpp Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE*/
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	/*"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE*/
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	/*"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA*/
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	/*"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303?discport=30304", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303?discport=30304",  // INFURA*/
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	/*"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",*/
}
