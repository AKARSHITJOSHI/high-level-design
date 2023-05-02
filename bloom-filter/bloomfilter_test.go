package filter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	bloomFilter := NewBloomFilter()
	bloomFilter.Add("Akarshit")
	bloomFilter.Add("Jitesh")
	bloomFilter.Add("Tarun")
	bloomFilter.Add("Aman")
}

func TestIsPresent(t *testing.T) {
	bloomFilter := NewBloomFilter()
	bloomFilter.Add("Akarshit")
	bloomFilter.Add("Jitesh")
	bloomFilter.Add("Tarun")
	bloomFilter.Add("Aman")
	require.Equal(t, bloomFilter.IsPresent("Akarshit"), true)
	require.Equal(t, bloomFilter.IsPresent("Aman"), true)
	require.Equal(t, bloomFilter.IsPresent("Harshit"), false)
	require.NotEqual(t, bloomFilter.IsPresent("Harshit"), true)
	bloomFilter.Clear()
	require.NotEqual(t, bloomFilter.IsPresent("Akarshit"), true)

}
