package UniName

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_brazil_full_name(t *testing.T) {
	name := WesternName{}
	name.Original = "Prof.ª João do Carmo Mão de    ferro e Cunha de almeida Santa rita Santos de Abreu JR."
	name.Nick = "Longa"
	name.Prefix = "Prof.ª"
	name.First = "João"
	name.FirstSocial = "Maria"
	name.BeforeMiddle = "do"
	name.Middle = "Carmo Mão     de Ferro e cunha de almeida Santa Rita Santos"
	name.BeforeLast = "de"
	name.Last = "abreu"
	name.Suffix = "JR."

	assert.Equal(t, "JOÃO DO CARMO MÃO DE FERRO E CUNHA DE ALMEIDA SANTA RITA SANTOS DE ABREU JÚNIOR", brazil_full_name(name, STYLE_BRAZIL_OFFICIAL))
	assert.Equal(t, "(Longa) Prof.ª Maria do Carmo Mão de Ferro e Cunha de Almeida Santa Rita Santos de Abreu Jr.", brazil_full_name(name, STYLE_BRAZIL_COMMON))
	assert.Equal(t, "(Longa) Prof.ª Maria do C M de F e C de A S R S de Abreu JR.", brazil_abrev(name, ABREV_CLEAN, ABREV_AS_MUCH_AS_RESONABLE))
	assert.Equal(t, "(Longa) Prof.ª Maria do C. M. de F. e C. de A. S. R. S. de Abreu JR.", brazil_abrev(name, ABREV_PERIOD, ABREV_AS_MUCH_AS_RESONABLE))
}

func Test_brazil_fill(t *testing.T) {
	name := brazil_fill("Prof.ª João do Carmo Mão de    ferro e Cunha de almeida Santa rita Santos de Abreu Neta")
	assert.Equal(t, "Longa", name.Nick)
	assert.Equal(t, "Prof.ª", name.Prefix)
	assert.Equal(t, "João", name.First)
	assert.Equal(t, "Maria", name.FirstSocial)
	assert.Equal(t, "do", name.BeforeMiddle)
	assert.Equal(t, "Carmo Mão de Ferro e Cunha de Almeida Santa Rita Santos", name.Middle)
	assert.Equal(t, "de", name.BeforeLast)
	assert.Equal(t, "Abreu", name.Last)
	assert.Equal(t, "Neta", name.Suffix)
}
