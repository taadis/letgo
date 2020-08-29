package faq100

type People struct {
	Name string
	Age  int
}

func Parse() {
	m := make(map[string]*People)
	ps := []People{
		{Name: "张三", Age: 3},
		{Name: "李四", Age: 4},
		{Name: "王五", Age: 5},
	}
	for _, p := range ps {
		m[p.Name] = &p
	}
}

// 以上代码有什么问题, 说明原因.
