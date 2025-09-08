---
frameworks:
- 其他
license: Apache License 2.0
tasks:
- keyword-spotting
---

sherpa 中自定义唤醒词模型，训练数据为 wenetspeech L （10000 小时）数据，模型大小约为 3.3 M， 建模单元为拼音（声母 + 韵母）。
使用 icefall 训练，已转换为 onnx 格式，此仓库主要是给 sherpa-onnx 这个 inference 引擎使用。

模型结构为 zipformer 模型， 本质是一个非常小的语音识别模型，为了实现唤醒词的功能，我们在解码端做了一些修改和约束。
支持自定义唤醒词，数量不限，效果需要单个调整参数。


#### Clone with HTTP
```bash
 git lfs install
 git clone https://www.modelscope.cn/pkufool/sherpa-onnx-kws-zipformer-wenetspeech-3.3M-2024-01-01.git
```


2.配置模型文件

模型包下载后包含以下文件:

```
sherpa-onnx-kws-zipformer-wenetspeech-3.3M-2024-01-01/
├── encoder-epoch-12-avg-2-chunk-16-left-64.int8.onnx    # 速度优先
├── encoder-epoch-12-avg-2-chunk-16-left-64.onnx         # 
├── encoder-epoch-99-avg-1-chunk-16-left-64.int8.onnx    # 速度优先 
├── encoder-epoch-99-avg-1-chunk-16-left-64.onnx         # 精度优先
├── decoder-epoch-12-avg-2-chunk-16-left-64.onnx         #
├── decoder-epoch-99-avg-1-chunk-16-left-64.onnx         # 精度优先
├── joiner-epoch-12-avg-2-chunk-16-left-64.int8.onnx     # 速度优先
├── joiner-epoch-12-avg-2-chunk-16-left-64.onnx          #
├── joiner-epoch-99-avg-1-chunk-16-left-64.int8.onnx     # 速度优先
├── joiner-epoch-99-avg-1-chunk-16-left-64.onnx          # 精度优先
├── tokens.txt                    # Token映射表（必需）
├── keywords_raw.txt              # 原始关键词（可选，用于生成）
├── keywords.txt                  # 现成的
├── test_wavs/                    # 测试音频（可选）
├── configuration.json            # 模型元信息（可选）
└── README.md                     # 说明文档（可选）

```

3.选择配置方案

方案一:精度优先(推荐)

```
cd sherpa-onnx-kws-zipformer-wenetspeech-3.3M-2024-01-01

# 复制精度优先的epoch-99 fp32三件套
cp encoder-epoch-99-avg-1-chunk-16-left-64.onnx ../models/encoder.onnx
cp decoder-epoch-99-avg-1-chunk-16-left-64.onnx ../models/decoder.onnx  
cp joiner-epoch-99-avg-1-chunk-16-left-64.onnx ../models/joiner.onnx

# 复制配套文件
cp tokens.txt ../models/tokens.txt
cp keywords_raw.txt ../models/keywords_raw.txt  # 可选

```

方案二:速度优先

```
cd sherpa-onnx-kws-zipformer-wenetspeech-3.3M-2024-01-01

# 复制速度优先的epoch-99 int8三件套
cp encoder-epoch-99-avg-1-chunk-16-left-64.int8.onnx ../models/encoder.onnx
cp decoder-epoch-99-avg-1-chunk-16-left-64.onnx ../models/decoder.onnx
cp joiner-epoch-99-avg-1-chunk-16-left-64.int8.onnx ../models/joiner.onnx

# 复制配套文件  
cp tokens.txt ../models/tokens.txt

```

### 注意事项

- **不要混用fp32与int8**:三个模型文件必须保持一致的精度
- **优先选择epoch-99**:比epoch-12训练更充分,精度更高
- **必需文件**:
    - `encoder.onnx`
    - `decoder.onnx`
    - `joiner.onnx`
    - `tokens.txt`
    - `keywords.txt`

## 最终模型文件结构

配置完成后,你的 models 目录下应该包含以下内容

```
models/
├── encoder.onnx      # 编码器模型（重命名后）
├── decoder.onnx      # 解码器模型（重命名后） 
├── joiner.onnx       # 连接器模型（重命名后）
├── tokens.txt        # 拼音Token映射表（228行版本）
├── keywords.txt      # 关键词配置文件（需创建）
└── keywords_raw.txt  # 原始关键词文件（可选）

```
