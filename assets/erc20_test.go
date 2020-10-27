/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */
package assets

import (
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/polynetwork/fabric_chaincode/utils"
	"github.com/polynetwork/poly/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	hdr0 = "00000000db056dd100000000000000000000000000000000000000000000000000000000000000000000000031398296a7e89188e0e99c6f46979470a8e59e5ca225181c13d003947fe0ea5a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e305f000000001dac2b7c00000000fdb2037b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303238313732393138353430623262353132656165313837326132613265336132386439383963363064393564616238383239616461376437646437303664363538227d2c7b22696e646578223a322c226964223a2231323035303338623861663632313065636664636263616232323535326566386438636634316336663836663963663961623533643836353734316366646238333366303662227d2c7b22696e646578223a332c226964223a2231323035303234383261636236353634623139623930363533663665396338303632393265386161383366373865376139333832613234613665666534316330633036663339227d2c7b22696e646578223a342c226964223a2231323035303236373939333061343261616633633639373938636138613366313265313334633031393430353831386437383364313137343865303339646538353135393838227d5d2c22706f735f7461626c65223a5b332c322c342c312c332c342c322c312c322c312c322c342c332c332c322c322c342c342c312c342c332c342c342c332c342c342c322c322c342c312c322c312c332c312c322c332c312c312c312c332c312c332c322c322c312c332c342c342c322c342c332c332c342c322c312c322c332c312c332c315d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d76776c36a8c30384672529690f733d95a58bbd940000"
	hdr60000 = "00000000db056dd100000000dad99a868e4bdb051a703b8992c6b2c1455d9e25bd3c3472ef85a8bc7c1a07d4bf67d4747270d9aadd54903398c155c9f6be7c9881d7bfe78ad096ccafbec5730000000000000000000000000000000000000000000000000000000000000000acb6a5aedc6498960e44ca226263a97f35da2187147407e1b639c99024f3dca379657c5f60ea00007df65806481c9877fda4037b226c6561646572223a342c227672665f76616c7565223a22424371676b554f2b4670505041617a7756645550745376794457704c34304d534f3753346e4147326d59384947364565596c4772334b3970434d662b65765a4f6654694a66636a3530314131316b7a446c686139322b593d222c227672665f70726f6f66223a22322f516954516e376350555868496e63416a63684c6d59484e6363754b3739397944443456717234465264425049616f39754d6f744962332f47463448445042643543466b2f6434456f78426f627232697337574b413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a36303030302c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a322c226e223a342c2263223a312c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a332c226964223a2231323035303234383261636236353634623139623930363533663665396338303632393265386161383366373865376139333832613234613665666534316330633036663339227d2c7b22696e646578223a322c226964223a2231323035303338623861663632313065636664636263616232323535326566386438636634316336663836663963663961623533643836353734316366646238333366303662227d2c7b22696e646578223a312c226964223a2231323035303238313732393138353430623262353132656165313837326132613265336132386439383963363064393564616238383239616461376437646437303664363538227d2c7b22696e646578223a342c226964223a2231323035303236373939333061343261616633633639373938636138613366313265313334633031393430353831386437383364313137343865303339646538353135393838227d5d2c22706f735f7461626c65223a5b332c322c332c332c322c322c342c312c342c312c312c342c312c332c312c312c342c322c322c322c332c342c332c342c322c342c322c342c342c342c342c332c322c332c342c322c342c332c332c322c312c312c322c342c312c332c312c312c332c322c332c312c322c332c312c312c342c312c332c325d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d76776c36a8c30384672529690f733d95a58bbd940323120502679930a42aaf3c69798ca8a3f12e134c019405818d783d11748e039de8515988231205028172918540b2b512eae1872a2a2e3a28d989c60d95dab8829ada7d7dd706d65823120502482acb6564b19b90653f6e9c806292e8aa83f78e7a9382a24a6efe41c0c06f390342011c2eefad8c4c77e3e431cc2fb3cc51a8109b67f449a4861a9725b3a321a035928a5ed6e9d15c02db4aef61afb2aa0b058e88367ef9debec020f15d5fe2f40ba52242011c1410f37cf9dd5aece1da6fca5f4285e234828ac100e69d1631e6e7cc742fe8951281ced9fc6efe4fc7aae5b1d52336de9fd533d5b4357e745f1703352f01bd8a42011bea7b76bd6afd197732a1ab08dd7ef9be8e7adbe83e42d8369639b926ec29ca4b3844cd13dd7e39ec307c9f6a16f3ff5447a8c7f9677e14a00f0494af56342e9c"

	hdr1 = "00000000db056dd100000000244bb1b9a01a2b4fd3fe6601b2d046b6cf695ba474cd0fe95f2cf4bd8108488e0000000000000000000000000000000000000000000000000000000000000000302b67ceb77b47c6b635ac3082894538b693e972568eae5c8733a29b90fe3bc6db80df9bddff8dc9f322ee7499c5744a47bf91913087d4cfdff1ab0106dd1608a08e865fb09b0100ad86f58eaf57f903fd10017b226c6561646572223a312c227672665f76616c7565223a224246724358307254384d7753536f6a6d5375585943754258796b686946303046776d653258726d565274695a5a386b674c356a3642324b703245534b64445871596c45634865462b33667347314c73724439736a70596b3d222c227672665f70726f6f66223a22724b635264306171796d333634305033462b6e36462f5447487a677745672b55473948772b666756623743717a504d6d7268715453696e6b4d4b4c51776d492b494f72496f52784e446a6176624d642f752b716464773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a36303030302c226e65775f636861696e5f636f6e666967223a6e756c6c7d000000000000000000000000000000000000000004231205028172918540b2b512eae1872a2a2e3a28d989c60d95dab8829ada7d7dd706d658231205038b8af6210ecfdcbcab22552ef8d8cf41c6f86f9cf9ab53d865741cfdb833f06b23120502679930a42aaf3c69798ca8a3f12e134c019405818d783d11748e039de851598823120502482acb6564b19b90653f6e9c806292e8aa83f78e7a9382a24a6efe41c0c06f390442011c4a549789967cc766bb9a60b2faae947bec19ee2e2d84845638789a3258586dc461d451d63b9a6afc29a436d5d4204eb695e2f800a9b3e59dc41677b240932d7742011bf808cf0daa407d3bef5889ac0a5071d7d1bde99a861a4e663dc82948a38279174446d7c18bafccb94274cafccad4e9530fce78d43a2f4fe06f977233de58b1ff42011b0c7c1bc4cc2eed0427a227dc74f6bd3264e16b6b5b9cbffb8ffd2f892266b1ff5ef9ab73533c0d8d9b6bf99ba04edd2d173163786dfc772fbb15ae9ddedc80fb42011c09d821760428588be4ecaabf99d66846577dc662240268354879de98a1e3a75216491e4c926b11ee68fbbca7844722f9ec618e70082b3bbe5906b55a974ed317"
	proof1 = "ef20fb6cc83379d4775fd8cbceaf824e8dc02f69e19c142f315580df7e2dcd24256b0200000000000000200000000000000000000000000000000000000000000000000000000000000001200bef4274080e792fd43427d5e3231b10179514455eee4a14fdbe9d2caba60600142eea349947f93c3b9b74fbcf141e102add510ece0600000000000000145411cbe06a0895d044670c5ab256bec76ae40c7c06756e6c6f636b4a14cc85571cd21ea6f66bbfb78121e2d2aa153031fb1434f00110bad3236f01468799d44fe04d7deb25f000e8764817000000000000000000000000000000000000000000000000000000"

	rootCA = `-----BEGIN CERTIFICATE-----
MIICKDCCAc+gAwIBAgIRAN4EisCV7Y+rbW2hHV7wI0wwCgYIKoZIzj0EAwIwczEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh
Lm9yZzEuZXhhbXBsZS5jb20wHhcNMjAxMDExMTg1NzAwWhcNMzAxMDA5MTg1NzAw
WjBqMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN
U2FuIEZyYW5jaXNjbzENMAsGA1UECxMEcGVlcjEfMB0GA1UEAxMWcGVlcjAub3Jn
MS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIlsw55yk3JX
yqtkpCrUsFK5X5wwcfaB3F2SggaW5PPTC0QWx3qIXLlPCK67bnX4w8fpG3ECE2qI
W3dJ9pFiN0KjTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAAMCsGA1Ud
IwQkMCKAIDdSh00xsy2nqjtFAK5YMYIrU5CrVLzVMJTuIqBRnftjMAoGCCqGSM49
BAMCA0cAMEQCIE6oFsTk+feM0FgPyzrAXz6X6T67Tx9t4EkZT/OoezD7AiBFElLQ
09lFFYvdtoQ/6rTc8TugxcWIlwgM4w6W9996+g==
-----END CERTIFICATE-----`

	newCA = `-----BEGIN CERTIFICATE-----
MIICHjCCAcWgAwIBAgIRAKU15UAdRc3gZQuCCdYE2SIwCgYIKoZIzj0EAwIwaTEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xFDASBgNVBAoTC2V4YW1wbGUuY29tMRcwFQYDVQQDEw5jYS5leGFt
cGxlLmNvbTAeFw0yMDEwMDkwMjQ5MDBaFw0zMDEwMDcwMjQ5MDBaMGoxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNp
c2NvMRAwDgYDVQQLEwdvcmRlcmVyMRwwGgYDVQQDExNvcmRlcmVyLmV4YW1wbGUu
Y29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEA67IcH48n8fpLoT9MjyDT6Qh
QZqGe5KXHG9sqHJdIbJoYpnHMxkletVrqI35Y6sgp4w9Sy+8jTvReHc1+fchwKNN
MEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYDVR0jBCQwIoAgfi+u
kqWiPFOtT8mCFDWk2Rbl5JDHW1dwJRmcEyihyqkwCgYIKoZIzj0EAwIDRwAwRAIg
HNzfr04Jzi4J/p1UZn1U14JM8S6ym65/BxmH9uqepM8CIA5/tfv6aZ53PpOVYsrs
zQW7eQxTo228awU1AIwsA95+
-----END CERTIFICATE-----`

	addr1 = "8b4dc42434c360d7999cff078be5f028dbb06484"
	addr2 = "ca8878697484174d2e7b1f874f78b446d6f22066"
)

func prepareEnv(isLp string) (*ERC20TokenImpl, *utils.CCStubMock) {
	impl := &ERC20TokenImpl{}
	mock := &utils.CCStubMock{
		CA: rootCA,
	}
	mock.Mem = make(map[string][]byte)
	mock.Args = [][]byte{
		[]byte("polyEth"),
		[]byte("pEth"),
		[]byte("18"),
		[]byte("1000000000000000000000000000"),
		[]byte("test"),
		[]byte(isLp),
	}
	resp := impl.Init(mock)
	if resp.Status != shim.OK {
		fmt.Println(resp.GetMessage())
	}

	return impl, mock
}

func TestERC20TokenImpl_name(t *testing.T) {
	impl, mock := prepareEnv("false")
	resp := impl.name(mock)
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, []byte("polyEth"), resp.Payload)
}

func TestERC20TokenImpl_decimal(t *testing.T) {
	impl, mock := prepareEnv("false")
	resp := impl.decimal(mock)

	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, big.NewInt(18).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_balanceOf(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte(addr1),
	}
	resp := impl.balanceOf(mock, mock.GetArgs())
	ts, _ := big.NewInt(0).SetString("1000000000000000000000000000", 10)

	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, ts.Bytes(), resp.Payload)
}

func TestERC20TokenImpl_mint(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte(addr2),
		[]byte("10000"),
	}
	resp := impl.mint(mock, mock.GetArgs())
	amt, _ := big.NewInt(0).SetString("10000", 10)

	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.balanceOf(mock, [][]byte {
		[]byte(addr2),
	})
	assert.Equal(t, amt.Bytes(), resp.Payload)

	resp = impl.totalSupply(mock)
	ts, _ := big.NewInt(0).SetString("1000000000000000000000000000", 10)
	assert.Equal(t, ts.Add(ts, amt).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_transfer(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte(addr2),
		[]byte("10000"),
	}
	resp := impl.transfer(mock, mock.GetArgs())
	amt, _ := big.NewInt(0).SetString("10000", 10)

	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.balanceOf(mock, [][]byte {
		[]byte(addr2),
	})
	assert.Equal(t, amt.Bytes(), resp.Payload)

	resp = impl.balanceOf(mock, [][]byte {
		[]byte(addr1),
	})
	ts, _ := big.NewInt(0).SetString("1000000000000000000000000000", 10)
	assert.Equal(t, ts.Sub(ts, amt).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_approve(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte(addr2),
		[]byte("10000"),
	}
	resp := impl.approve(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.allowance(mock, [][]byte {
		[]byte(addr1),
		[]byte(addr2),
	})
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, big.NewInt(10000).Bytes(), resp.Payload)

	mock.SetNewCA(newCA)
	resp = impl.transferFrom(mock, [][]byte{
		[]byte(addr1),
		[]byte(addr2),
		[]byte("10000"),
	})
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.balanceOf(mock, [][]byte {
		[]byte(addr2),
	})
	assert.Equal(t, big.NewInt(10000).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_decreaseAllowance(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte(addr2),
		[]byte("10000"),
	}
	resp := impl.approve(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.decreaseAllowance(mock, [][]byte{
		[]byte(addr2),
		[]byte("1000"),
	})
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.allowance(mock, [][]byte {
		[]byte(addr1),
		[]byte(addr2),
	})
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, big.NewInt(9000).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_burn(t *testing.T) {
	impl, mock := prepareEnv("false")
	mock.Args = [][]byte {
		[]byte("10000"),
	}
	resp := impl.burn(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.balanceOf(mock, [][]byte {
		[]byte(addr1),
	})
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, big.NewInt(9000).Bytes(), resp.Payload)
}

func TestERC20TokenImpl_bindProxyHash(t *testing.T) {
	impl, mock := prepareEnv("true")
	mock.Args = [][]byte {
		[]byte("2"),
		[]byte(addr1),
	}
	resp := impl.bindProxyHash(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.getProxyHash(mock, [][]byte {
		[]byte("2"),
	})
	raw, _ := hex.DecodeString(addr1)
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, raw, resp.Payload)
}

func TestERC20TokenImpl_bindAssetHash(t *testing.T) {
	impl, mock := prepareEnv("true")
	mock.Args = [][]byte {
		[]byte("2"),
		[]byte(addr1),
	}

	resp := impl.bindAssetHash(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")

	resp = impl.getAssetHash(mock, [][]byte {
		[]byte("2"),
	})
	raw, _ := hex.DecodeString(addr1)
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
	assert.Equal(t, raw, resp.Payload)
}

func TestERC20TokenImpl_lock(t *testing.T) {
	impl, mock := prepareEnv("true")
	mock.Args = [][]byte {
		[]byte("ccm"),
	}
	resp := impl.setManager(mock, mock.GetArgs())
	mock.Args = [][]byte {
		[]byte("2"),
		[]byte(addr1),
	}
	resp = impl.bindAssetHash(mock, mock.GetArgs())
	mock.Args = [][]byte {
		[]byte("2"),
		[]byte(addr1),
	}
	resp = impl.bindProxyHash(mock, mock.GetArgs())

	to, _ := hex.DecodeString(addr1)
	resp = impl.getLockProxyAddr(mock)
	amt := big.NewInt(1000)
	impl.transferLogic(mock, resp.Payload, to, amt)

	mock.Args = [][]byte {
		[]byte("2"),
		[]byte(addr1),
		[]byte(amt.String()),
	}
	resp = impl.lock(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
}

func TestERC20TokenImpl_unlock(t *testing.T) {
	impl, mock := prepareEnv("true")
	mock.Args = [][]byte {
		[]byte("ccm"),
	}
	resp := impl.setManager(mock, mock.GetArgs())

	to, _ := hex.DecodeString(addr1)
	txArgs := &TxArgs{
		Amount: 1000,
		ToAddress: to,
		ToAssetHash: []byte("test"),
	}
	sink := common.NewZeroCopySink(nil)
	txArgs.Serialization(sink)

	mock.Args = [][]byte {
		[]byte(hex.EncodeToString(sink.Bytes())),
	}
	resp = impl.unlock(mock, mock.GetArgs())
	assert.Equal(t, true, shim.OK == resp.Status,  "wrong result")
}