package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// 限制大小的数组
	arr := New(3)

	// 测试能否正常添加
	err := arr.Add(1)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{1})

	// 测试超出size大小报错
	err = arr.Add(2)
	assert.Nil(t, err)
	err = arr.Add(3)
	assert.Nil(t, err)
	err = arr.Add(4)
	assert.ErrorContains(t, err, "the current data volume has reached the maximum")
	assert.Equal(t, arr.Data(), []int{1, 2, 3})

	// 不限制大小的数组
	arr = New()
	assert.Nil(t, err)
	err = arr.Add(1)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{1})
}

func TestData(t *testing.T) {
	arr := New(3)

	// 测试获取数组结果正确
	err := arr.Add(1)
	assert.Nil(t, err)
	err = arr.Add(2)
	assert.Nil(t, err)
	err = arr.Add(3)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{1, 2, 3})
}

func TestInsert(t *testing.T) {
	// 限制大小的数组
	arr := New(4)

	// 测试能否正常添加
	err := arr.Insert(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{1})

	// 插入头部位置
	err = arr.Insert(0, 2)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{2, 1})

	// 插入中间位置
	err = arr.Insert(1, 3)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{2, 3, 1})

	// 插入尾部位置
	err = arr.Insert(3, 4)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{2, 3, 1, 4})

	// 测试超出size大小报错
	err = arr.Insert(0, 1)
	assert.ErrorContains(t, err, "the current data volume has reached the maximum")
	assert.Equal(t, arr.Data(), []int{2, 3, 1, 4})

	// 测试插入位置小于0
	// 删除最后一个元素
	_, err = arr.Pop()
	assert.Nil(t, err)
	err = arr.Insert(-1, 1)
	assert.ErrorContains(t, err, "index must be between")
	assert.Equal(t, arr.Data(), []int{2, 3, 1})

	// 测试插入位置大于数组长度
	err = arr.Insert(5, 1)
	assert.ErrorContains(t, err, "index must be between")
	assert.Equal(t, arr.Data(), []int{2, 3, 1})

	// 不限制大小的数组
	arr = New()
	assert.Nil(t, err)
	err = arr.Add(1)
	assert.Nil(t, err)
	assert.Equal(t, arr.Data(), []int{1})
}
