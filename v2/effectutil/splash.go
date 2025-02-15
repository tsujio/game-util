package effectutil

import (
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SplashParticle struct {
	x, y   float64
	vx, vy float64
	theta  float64
	omega  float64
	color  color.Color
	size   float64
}

type SplashEffect struct {
	ticks, maxTicks uint64
	particles       []SplashParticle
	opts            SplashEffectOptions
}

type SplashEffectOptions struct {
	Count              int
	Color              color.Color
	Size               float64
	AngularVelocity    float64
	AngleMin, AngleMax float64
	Speed              float64
	Ay                 float64
	Random             *rand.Rand
}

func NewSplashEffect(x, y float64, span uint64, opts *SplashEffectOptions) *SplashEffect {
	random := opts.Random
	if random == nil {
		random = rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Unix())))
	}

	particles := make([]SplashParticle, opts.Count)
	for i := 0; i < opts.Count; i++ {
		p := &particles[i]
		p.x = x
		p.y = y
		p.color = opts.Color
		p.size = opts.Size

		angleMax := opts.AngleMax
		if opts.AngleMax == 0 && opts.AngleMin == 0 {
			angleMax = math.Pi * 2
		}
		theta := random.Float64()*(angleMax-opts.AngleMin) + opts.AngleMin
		p.vx = opts.Speed * math.Cos(theta)
		p.vy = opts.Speed * math.Sin(theta)

		p.omega = opts.AngularVelocity
		if p.omega != 0 {
			p.theta = random.Float64() * math.Pi * 2
		}
	}

	e := &SplashEffect{
		maxTicks:  span,
		particles: particles,
		opts:      *opts,
	}

	return e
}

func (e *SplashEffect) Type() string {
	return "splash"
}

func (e *SplashEffect) Update() {
	e.ticks++

	for i := range e.particles {
		p := &e.particles[i]
		p.x += p.vx
		p.y += p.vy

		p.vy += e.opts.Ay

		p.theta += p.omega
	}
}

var splashParticleImg = ebiten.NewImage(1, 1)

func (e *SplashEffect) Draw(dst *ebiten.Image) {
	for i := range e.particles {
		p := &e.particles[i]
		vector.DrawFilledRect(splashParticleImg, 0, 0, 1, 1, p.color, false)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(-0.5, -0.5)
		opts.GeoM.Scale(p.size, p.size)
		opts.GeoM.Rotate(p.theta)
		opts.GeoM.Translate(p.x, p.y)
		dst.DrawImage(splashParticleImg, opts)
	}
}

func (e *SplashEffect) Finished() bool {
	return e.ticks > e.maxTicks
}
