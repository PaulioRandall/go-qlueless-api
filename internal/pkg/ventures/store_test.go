package ventures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVentureStore_GetAll_1(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		VentureID:   "1",
		Description: "1",
		State:       "1",
	}
	b := Venture{
		VentureID:   "2",
		Description: "2",
		State:       "2",
	}

	store.items["1"] = a
	store.items["2"] = b

	s := store.GetAll()

	require.Len(t, s, 2)
	assert.Contains(t, s, a)
	assert.Contains(t, s, b)
}

func TestVentureStore_GetAll_2(t *testing.T) {
	store := NewVentureStore()
	s := store.GetAll()
	require.Empty(t, s)
}

func TestVentureStore_GetAllAlive_1(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		VentureID:   "1",
		Description: "1",
		State:       "1",
		IsAlive:     true,
	}
	b := Venture{
		VentureID:   "2",
		Description: "2",
		State:       "2",
		IsAlive:     false,
	}

	store.items["1"] = a
	store.items["2"] = b

	s := store.GetAllAlive()

	require.Len(t, s, 1)
	assert.Contains(t, s, a)
}

func TestVentureStore_GetAllAlive_2(t *testing.T) {
	store := NewVentureStore()
	s := store.GetAllAlive()
	require.Empty(t, s)
}

func TestVentureStore_Get_1(t *testing.T) {
	store := NewVentureStore()
	aIn := Venture{
		VentureID:   "1",
		Description: "1",
		State:       "1",
	}
	bIn := Venture{
		VentureID:   "2",
		Description: "2",
		State:       "2",
	}

	store.items["1"] = aIn
	store.items["2"] = bIn

	aOut, ok := store.Get("1")
	require.True(t, ok)
	assert.Equal(t, aIn, aOut)

	bOut, ok := store.Get("2")
	require.True(t, ok)
	assert.Equal(t, bIn, bOut)
}

func TestVentureStore_Get_2(t *testing.T) {
	store := NewVentureStore()
	_, ok := store.Get("1")
	require.False(t, ok)
}

func TestVentureStore_Get_3(t *testing.T) {
	store := NewVentureStore()
	aIn := Venture{
		VentureID:   "1",
		Description: "1",
		State:       "1",
	}
	bIn := Venture{
		VentureID:   "2",
		Description: "2",
		State:       "2",
	}

	store.items["1"] = aIn
	store.items["2"] = bIn

	_, ok := store.Get("3")
	require.False(t, ok)
}
