package day1_test

import (
	"testing"

	"github.com/geisonbiazus/advent_of_code/2020/day1"
	"github.com/stretchr/testify/assert"
)

func TestFindEntries(t *testing.T) {
	t.Run("It returns empty when there is no entry that sums to 2020", func(t *testing.T) {
		assert.Equal(t, []int{}, day1.FindEntriesThatSumTo2020(2, []int{}))
		assert.Equal(t, []int{}, day1.FindEntriesThatSumTo2020(2, []int{2}))
		assert.Equal(t, []int{}, day1.FindEntriesThatSumTo2020(2, []int{2, 3}))
		assert.Equal(t, []int{}, day1.FindEntriesThatSumTo2020(2, []int{2020, 3}))
	})

	t.Run("It returns the two entries that sums to 2020", func(t *testing.T) {
		assert.Equal(t, []int{2020, 0}, day1.FindEntriesThatSumTo2020(2, []int{2020, 0}))
		assert.Equal(t, []int{1, 2019}, day1.FindEntriesThatSumTo2020(2, []int{1, 2019}))
		assert.Equal(t, []int{1, 2019}, day1.FindEntriesThatSumTo2020(2, []int{1, 100, 2019}))
		assert.Equal(t, []int{1, 2019}, day1.FindEntriesThatSumTo2020(2, []int{100, 1, 2019}))
	})

	t.Run("It finds the given number of entries that sums to 2020", func(t *testing.T) {
		assert.Equal(t, []int{2020}, day1.FindEntriesThatSumTo2020(1, []int{2020, 0}))
		assert.Equal(t, []int{1000, 1001, 19}, day1.FindEntriesThatSumTo2020(3, []int{1000, 1001, 18, 19}))
	})
}

func TestSolvePuzzle(t *testing.T) {
	t.Run("It returns 0 when the two numbers cannot be found", func(t *testing.T) {
		assert.Equal(t, 0, day1.SolvePuzzle(2, []int{}))
		assert.Equal(t, 0, day1.SolvePuzzle(2, []int{2}))
		assert.Equal(t, 0, day1.SolvePuzzle(2, []int{2, 2, 4}))
	})

	t.Run("It returns the two numbers that sums to 2020 multiplied", func(t *testing.T) {
		assert.Equal(t, 2019, day1.SolvePuzzle(2, []int{2019, 1}))
		assert.Equal(t, 4036, day1.SolvePuzzle(2, []int{2018, 2}))
	})

	t.Run("It returns the 3 numbers that sums to 2020 multiplied", func(t *testing.T) {
		assert.Equal(t, 2018, day1.SolvePuzzle(3, []int{2018, 1, 1}))
		assert.Equal(t, 4034, day1.SolvePuzzle(3, []int{2017, 2, 90, 1}))
	})

	t.Run("It solves the puzzle", func(t *testing.T) {
		assert.Equal(t, 776064, day1.SolvePuzzle(2, day1Input))
		assert.Equal(t, 6964490, day1.SolvePuzzle(3, day1Input))
	})
}

var day1Input = []int{
	1934, 1702, 1571, 1737, 1977, 1531, 1428, 1695, 1794, 1101, 13, 1164, 1235, 1289, 1736, 1814, 1363, 1147, 1111, 1431, 1765, 1515, 1184, 1036, 1803, 1791, 1638, 1809, 1283, 1980, 1854, 1878, 1574, 1352, 1151, 730, 1581, 1990, 1919, 2003, 1538, 1663, 1735, 1772, 1830, 1152, 1022, 1774, 1544, 1551, 1835, 1383, 1614, 1396, 1715, 1530, 295, 1208, 1978, 1104, 1691, 1176, 1183, 1909, 1192, 1535, 1924, 1268, 1969, 1954, 1760, 1077, 1734, 1371, 1676, 1933, 1400, 1928, 1982, 1541, 1106, 1248, 1346, 1782, 1142, 1849, 1798, 1362, 1379, 1886, 1265, 1226, 1751, 1575, 1027, 1710, 1601, 1205, 1922, 1452, 1206, 1263, 2000, 1957, 1951, 1834, 1533, 1149, 1245, 1564, 1182, 1237, 1013, 1254, 1895, 1504, 1480, 1556, 1821, 1589, 1864, 1573, 1698, 1927, 1434, 516, 1722, 1360, 1940, 1212, 1329, 1675, 1812, 1917, 1302, 1604, 1336, 1233, 1405, 1179, 1169, 1081, 1941, 1553, 1236, 1824, 1923, 1938, 1475, 1446, 1545, 1853, 1664, 317, 1489, 1884, 1743, 1621, 1128, 1474, 1505, 394, 1387, 1509, 1627, 1914, 1913, 1949, 1843, 1847, 1882, 1486, 1082, 1802, 1645, 1690, 1629, 1377, 2004, 1044, 1191, 1014, 1857, 1813, 1572, 1055, 1002, 1721, 1273, 1417, 1968, 1888, 1863, 1278, 1141, 1964, 1259, 1823, 1181, 1779,
}
