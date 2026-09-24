package api

import "testing"

func TestParseSimpleGres(t *testing.T) {
	node := NodeData{}
	gres := "gpu:gtx_1080:1"

	node.SetNodeGPUType(&gres)

	if len(node.GpuType) != 1 {
		t.Error("Expected one GPU type")
	}
	if node.GpuType[0] != "gtx_1080" {
		t.Errorf("Incorrect GPU type = %v. wanted gtx_1080", node.GpuType[0])
	}
}

func TestParseGresSocketIndex(t *testing.T) {
	node := NodeData{}
	gres := "gpu:h200:8(S:0-1)"

	node.SetNodeGPUType(&gres)

	if len(node.GpuType) != 1 {
		t.Error("Expected one GPU type")
	}
	if node.GpuType[0] != "h200" {
		t.Errorf("Incorrect GPU type = %v. wanted h200", node.GpuType[0])
	}
}

func TestParseMultiGpuNodeGres(t *testing.T) {
	node := NodeData{}
	gres := "gpu:a100_80gb_pcie_3g.40gb:2(S:0-1),gpu:a100_80gb_pcie_2g.20gb:4(S:0-1)"

	node.SetNodeGPUType(&gres)

	if len(node.GpuType) != 2 {
		t.Error("Expected two GPU types")
	}
	if node.GpuType[0] != "a100_80gb_pcie_3g.40gb" {
		t.Errorf("Incorrect GPU type = %v. wanted a100_80gb_pcie_3g.40gb", node.GpuType[0])
	}
	if node.GpuType[1] != "a100_80gb_pcie_2g.20gb" {
		t.Errorf("Incorrect GPU type = %v. wanted a100_80gb_pcie_2g.20gb", node.GpuType[1])
	}
}

func TestParseMultiGpuNodeTres(t *testing.T) {
	node := NodeData{}
	tres := "cpu=256,mem=512G,billing=1556,gres/gpu=6,gres/gpu:a100_80gb_pcie_2g.20gb=4,gres/gpu:a100_80gb_pcie_3g.40gb=2"
	empty := ""

	err := node.SetNodeGPUAllocated(&tres, &empty)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(node.GPUAllocated) != 2 {
		t.Error("Expected two GPU types")
	}
	if node.GPUAllocated["a100_80gb_pcie_3g.40gb"] != 2 {
		t.Error("Expected two 40gb GPUs")
	}
	if node.GPUAllocated["a100_80gb_pcie_2g.20gb"] != 4 {
		t.Error("Expected four 20gb GPUs")
	}
}
