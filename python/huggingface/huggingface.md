# clash

```bash
wget -O config.yaml  订阅地址 

wget -O Country.mmdb https://www.sub-speeder.com/client-download/Country.mmdb

./clash -d .

# 配置环境变量
export http_proxy="http://127.0.0.1:7890"
export https_proxy="http://127.0.0.1:7890"
export ftp_proxy="http://127.0.0.1:7890"
export no_proxy="localhost,127.0.0.1,::1"
```

# Huggingface

## 环境安装

```bash
# 安装pip
pip install --upgrade huggingface_hub

# 登录
huggingface-cli login
huggingface-cli logout

# 下载模型
huggingface-cli download --resume-download [model_nane] --local-dir [path]

# 配置缓存路径
export HF_DATASETS_CACHE="/root/autodl-tmp/cache/"
export HF_HOME="/root/autodl-tmp/cache/"
export HUGGINGFACE_HUB_CACHE="/root/autodl-tmp/cache/"
export TRANSFORMERS_CACHE="/root/autodl-tmp/cache/"
```

## 微调

```bash
CUDA_VISIBLE_DEVICES=0 python src/train_bash.py \
    --stage sft \
    --do_train \
    --model_name_or_path /root/autodl-tmp/llama3 \
    --dataset identity \
    --template default \
    --finetuning_type lora \
    --lora_target q_proj,v_proj \
    --output_dir /root/autodl-tmp/output \
    --overwrite_cache \
    --per_device_train_batch_size 4 \
    --gradient_accumulation_steps 4 \
    --lr_scheduler_type cosine \
    --logging_steps 10 \
    --save_steps 1000 \
    --learning_rate 5e-5 \
    --num_train_epochs 20.0 \
    --plot_loss \
    --fp16
```

## 导出

```bash
python src/export_model.py \
    --model_name_or_path meta-llama/Meta-Llama-3-8B-Instruct \
    --adapter_name_or_path /root/autodl-tmp/output \
    --template default \
    --finetuning_type lora \
    --export_dir /root/autodl-tmp/export_models \
    --export_size 2 \
    --export_legacy_format False
    
    
    
python src/export_model.py \
    --model_name_or_path /root/autodl-tmp/llama3 \
    --adapter_name_or_path /root/autodl-tmp/output \
    --template default \
    --finetuning_type lora \
    --export_dir /root/autodl-tmp/export_models \
    --export_size 2 \
    --export_legacy_format False
```

## 测试

```bash
python src/web_demo.py \
    --model_name_or_path /root/autodl-tmp/llama3 \
    --adapter_name_or_path /root/autodl-tmp/output \
    --template default \
    --finetuning_type lora
```

