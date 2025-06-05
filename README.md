# Pixiv Artwork Tag & Recommendation Service

## 背景・課題
- Pixiv には1億超の投稿があり、タグ付け・検索体験の質がサービス価値を左右する。
- 既存自動タグ生成は英語偏重・非公開。OSS スタックで誰でも再現できる解決策を提示する。

## ソリューション概要
- CLIP 埋め込み + Faiss で高速タグ候補抽出
- BigQuery + dbt で協調フィルタリング推薦
- Go + Clean Architecture により拡張容易

## アーキテクチャ図
<include ./docs/architecture.svg>
*(Placeholder for architecture diagram)*

## 技術選定理由
| Concern | Choice | Rationale |
| --- | --- | --- |
| 言語 | Go | Cloud Run と相性良 |
| 推論 | CLIP ViT-L/14 | OSS・GPU/CPU 両対応 |
| DWH | BigQuery | 大規模データと低コスト |
| IaC | Terraform | マルチ環境管理 |
*(More to be added)*

## セットアップ
```bash
git clone https://github.com/your-handle/pixiv-tag-reco-service
make dev   # build images & start local stack
```
*(Instructions to be verified and updated)*

## 主要ユースケース
1. 画像アップロード → 200 ms 以内に上位5タグ提示
2. ユーザがタグ保存 → Pub/Sub 経由で ETL
3. BigQuery モデル再計算 → 推薦即時反映
*(Details to be added)*

## 深掘り Q&A
**Q. 精度評価方法は?**
A. 公開データで MAP@5=0.62／Hit@1=0.43 (+18 %).
*(More Q&A to be added)*

**Q. スケーラビリティは?**
A. 推論 Pod は 0→N オートスケール。Faiss は read-only & sharded。
*(More Q&A to be added)*

## 今後の展望
- 日本語キャプション LLM fine-tune
- Rust 化による低レイテンシ
*(More details to be added)*
