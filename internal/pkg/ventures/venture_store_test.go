package ventures

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// ****************************************************************************
// VentureStore.GetAll()
// ****************************************************************************

func TestVentureStore_GetAll_1(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		ID:          "1",
		Description: "1",
		State:       "1",
	}
	b := Venture{
		ID:          "2",
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

// ****************************************************************************
// VentureStore.GetAllAlive()
// ****************************************************************************

func TestVentureStore_GetAllAlive_1(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		ID:          "1",
		Description: "1",
		State:       "1",
		IsAlive:     true,
	}
	b := Venture{
		ID:          "2",
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

// ****************************************************************************
// VentureStore.Get()
// ****************************************************************************

func TestVentureStore_Get_1(t *testing.T) {
	store := NewVentureStore()
	aIn := Venture{
		ID:          "1",
		Description: "1",
		State:       "1",
	}
	bIn := Venture{
		ID:          "2",
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
		ID:          "1",
		Description: "1",
		State:       "1",
	}
	bIn := Venture{
		ID:          "2",
		Description: "2",
		State:       "2",
	}

	store.items["1"] = aIn
	store.items["2"] = bIn

	_, ok := store.Get("3")
	require.False(t, ok)
}

// ****************************************************************************
// VentureStore.Add()
// ****************************************************************************

func TestVentureStor_Add_1(t *testing.T) {
	store := NewVentureStore()
	aIn := NewVenture{
		Description: "description",
		State:       "state",
	}

	aOut := store.Add(aIn)
	assert.Len(t, store.items, 1)
	assert.NotEmpty(t, aOut.ID)
	assert.Equal(t, "description", aOut.Description)
	assert.Equal(t, "state", aOut.State)
}

func TestVentureStor_Add_2(t *testing.T) {
	store := NewVentureStore()
	aIn := NewVenture{
		Description: "description",
		State:       "state",
	}
	bIn := NewVenture{
		Description: "description",
		State:       "state",
	}

	aOut := store.Add(aIn)
	bOut := store.Add(bIn)
	assert.Len(t, store.items, 2)
	assert.NotEmpty(t, aOut.ID)
	assert.NotEmpty(t, bOut.ID)
	assert.NotEqual(t, aOut.ID, bOut.ID)
	assert.Equal(t, "description", aOut.Description)
	assert.Equal(t, "description", bOut.Description)
	assert.Equal(t, "state", aOut.State)
	assert.Equal(t, "state", bOut.State)
}

// ****************************************************************************
// VentureStore.Update()
// ****************************************************************************

func TestVentureStore_Update_1(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		ID: "1",
	}
	b := Venture{
		ID: "2",
	}
	store.items["1"] = a
	store.items["2"] = b

	u := ModVenture{
		IDs:   "1,2",
		Props: "description,order_ids,state,is_alive,extra",
		Values: Venture{
			Description: "new",
			OrderIDs:    "new",
			State:       "new",
			IsAlive:     true,
			Extra:       "new",
		},
	}

	v := store.Update(&u)
	require.Len(t, v, 2)

	u.Values.ID = "1"
	u.Values.LastModified = v[0].LastModified
	assert.Equal(t, u.Values, v[0])
	aOut, ok := store.items["1"]
	require.True(t, ok)
	assert.Equal(t, u.Values, aOut)

	u.Values.ID = "2"
	assert.Equal(t, u.Values, v[1])
	bOut, ok := store.items["2"]
	require.True(t, ok)
	assert.Equal(t, u.Values, bOut)
}

func TestVentureStore_Update_2(t *testing.T) {
	store := NewVentureStore()
	au := ModVenture{
		IDs:   "1",
		Props: "description",
		Values: Venture{
			Description: "original",
		},
	}

	v := store.Update(&au)
	require.Empty(t, v)
}

func TestVentureStore_Update_3(t *testing.T) {
	store := NewVentureStore()
	a := Venture{
		ID: "1",
	}
	b := Venture{
		ID: "2",
	}
	store.items["1"] = a
	store.items["2"] = b

	u := ModVenture{
		IDs:   "1,2",
		Props: "IGNORED,SKIPPED",
		Values: Venture{
			Description: "new",
			OrderIDs:    "new",
			State:       "new",
			IsAlive:     true,
			Extra:       "new",
		},
	}

	v := store.Update(&u)
	require.Len(t, v, 2)

	a.LastModified = v[0].LastModified
	assert.Equal(t, a, v[0])
	b.LastModified = v[1].LastModified
	assert.Equal(t, b, v[1])

	aOut, ok := store.items["1"]
	require.True(t, ok)
	assert.Equal(t, a, aOut)

	bOut, ok := store.items["2"]
	require.True(t, ok)
	assert.Equal(t, b, bOut)
}
