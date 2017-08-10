package main

//#include<relic.h>
import "C"
import (
	"crypto/cipher"
	"crypto/sha256"
	"hash"

	kyber "gopkg.in/dedis/kyber.v1"
	"gopkg.in/dedis/kyber.v1/cipher/sha3"
)

type PairingGroup interface {
	kyber.Group // Standard Group operations

	PointGT() PointGT // Create new pairing-capable Point
}

type PointGT interface {
	kyber.Point // Standard Point operations

	// Compute the pairing of two points p1 and p2,
	// which must be in the associated groups G1 and G2 respectively.
	Pairing(p1, p2 kyber.Point) kyber.Point
}

type PairingSuite interface {
	G1() kyber.Group
	G2() kyber.Group
	GT() PairingGroup
}

type g1group struct {
	common
	curve int
}
type g2group struct {
	common
	curve int
}
type gtgroup struct {
	common
	p *Pairing
}

type Pairing struct {
	curve int
	g1    g1group
	g2    g2group
	gt    gtgroup
}

func NewPairing(curve int) *Pairing {
	ok := curve == CurveFp254BNb || curve == CurveFp382_1 || curve == CurveFp382_2
	if !ok {
		panic("pairing: unsupported curve")
	}
	C.ep_param_set(C.int(curve))
	p := &Pairing{curve: curve}
	p.g1.curve = curve
	p.g2.curve = curve
	p.gt.p = p
	return p
}

func NewPairingFp254BNb() *Pairing {
	return NewPairing(CurveFp254BNb)
}

func NewPairingFp382_1() *Pairing {
	return NewPairing(CurveFp382_1)
}

func NewPairingFp382_2() *Pairing {
	//TODO: Check what's the appropriate curve and modify misc.go.
	// Currently set to 0
	return NewPairing(CurveFp382_2)
}

func (p *Pairing) G1() kyber.Group {
	return &p.g1
}

func (p *Pairing) G2() kyber.Group {
	return &p.g2
}

func (p *Pairing) GT() PairingGroup {
	return &p.gt
}

func (p *Pairing) Cipher(key []byte, options ...interface{}) kyber.Cipher {
	return sha3.NewShakeCipher128(key, options...)
}

func (p *Pairing) Hash() hash.Hash {
	return sha256.New()
}
func (g *g1group) String() string {
	return curveName(g.curve) + "_G1"
}

func (g *g1group) PointLen() int {
	return g.Point().MarshalSize()
}

func (g *g1group) Point() kyber.Point {
	return newPointG1(generator(g.curve, 0))
}

func (g *g1group) NewKey(rand cipher.Stream) kyber.Scalar {
	return NewScalar().Pick(rand)
}

func (g *g2group) String() string {
	return curveName(g.curve) + "_G2"
}

func (g *g2group) PointLen() int {
	return g.Point().MarshalSize()
}

func (g *g2group) Point() kyber.Point {
	return newPointG2(generator(g.curve, 1))
}

func (g *g2group) NewKey(rand cipher.Stream) kyber.Scalar {
	return NewScalar().Pick(rand)
}

func (g *gtgroup) String() string {
	return curveName(g.p.curve) + "_GT"
}

func (g *gtgroup) PointLen() int {
	return g.Point().MarshalSize()
}

func (g *gtgroup) Point() kyber.Point {
	panic("PoinGT not properly imlemented")
	//XXX
}

func (g *gtgroup) PointGT() PointGT {
	return g.Point().(PointGT)
}

func (g *gtgroup) NewKey(rand cipher.Stream) kyber.Scalar {
	return NewScalar().Pick(rand)
}

type common struct{}

func (c *common) Hash() hash.Hash {
	return sha256.New()
}

func (c *common) Cipher(key []byte, options ...interface{}) kyber.Cipher {
	return sha3.NewShakeCipher128(key, options...)
}

func (c *common) PrimeOrder() bool {
	return true
}

func (c *common) ScalarLen() int {
	return int(C.RELIC_BN_BYTES)
}

func (c *common) Scalar() kyber.Scalar {
	return NewScalar()
}

func curveName(curve int) string {
	switch curve {
	case CurveFp254BNb:
		return "Fp254Nb"
	case CurveFp382_1:
		return "Fp382_1"
	case CurveFp382_2:
		return "Fp382_2"
	default:
		panic("pairing curve unknown")
	}
}

func generator(curve, group int) string {
	var gens [2]string
	switch curve {
	case CurveFp254BNb:
		gens[0] = Fp254_G1_Base_Str
		gens[1] = Fp254_G2_Base_Str
	case CurveFp382_1:
		gens[0] = Fp382_1_G1_Base_Str
		gens[1] = Fp382_1_G2_Base_Str
	case CurveFp382_2:
		gens[0] = Fp382_2_G1_Base_Str
		gens[1] = Fp382_2_G2_Base_Str
	default:
		panic("pairing curve unknown")
	}
	return gens[group]
}
