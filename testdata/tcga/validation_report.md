# TCGA Validation Report

Generated: 2026-03-03 16:05 UTC  
GENCODE transcripts loaded: 254070  
Cache load time: 43.358s  
Workers: 4 (GOMAXPROCS)

## Match Rates

| Study | Variants | Conseq Match | Conseq Mismatch | Conseq Rate | HGVSp Match | HGVSp Mismatch | HGVSp Rate | HGVSc Match | HGVSc Mismatch | HGVSc Rate |
|-------|----------|-------------|-----------------|-------------|-------------|----------------|------------|-------------|----------------|------------|
| blca_tcga_gdc | 115850 | 115649 | 1 | 99.8% | 109994 | 0 | 94.9% | 114177 | 0 | 98.6% |
| brca_tcga_gdc | 89012 | 88847 | 1 | 99.8% | 84167 | 0 | 94.6% | 87775 | 0 | 98.6% |
| chol_tcga_gdc | 3764 | 3757 | 0 | 99.8% | 3484 | 0 | 92.6% | 3695 | 0 | 98.2% |
| coad_tcga_gdc | 244552 | 244087 | 2 | 99.8% | 229187 | 0 | 93.7% | 240971 | 0 | 98.5% |
| gbm_tcga_gdc | 54870 | 54747 | 0 | 99.8% | 51900 | 0 | 94.6% | 54126 | 0 | 98.6% |
| luad_tcga_gdc | 190868 | 190445 | 1 | 99.8% | 181724 | 0 | 95.2% | 188218 | 0 | 98.6% |
| skcm_tcga_gdc | 353450 | 352554 | 3 | 99.7% | 336733 | 0 | 95.3% | 348715 | 0 | 98.7% |
| **Total** | **1052366** | **1050086** | **8** | **99.8%** | **997189** | **0** | **94.8%** | **1037677** | **0** | **98.6%** |

## Consequence Category Breakdown

| Study | delins_normalized | gene_model_change | match | mismatch | no_cds_data | position_shift | transcript_model_change | upstream_reclassified |
|-------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 1 | 0 | 114750 | 1 | 74 | 42 | 83 | 899 |
| brca_tcga_gdc | 2 | 0 | 88303 | 1 | 78 | 32 | 52 | 544 |
| chol_tcga_gdc | 0 | 0 | 3726 | 0 | 4 | 1 | 2 | 31 |
| coad_tcga_gdc | 8 | 2 | 242406 | 2 | 193 | 99 | 161 | 1681 |
| gbm_tcga_gdc | 0 | 0 | 54376 | 0 | 58 | 33 | 32 | 371 |
| luad_tcga_gdc | 1 | 0 | 189094 | 1 | 219 | 96 | 106 | 1351 |
| skcm_tcga_gdc | 5 | 3 | 349986 | 3 | 398 | 266 | 221 | 2568 |
| **Total** | **17** | **5** | **1042641** | **8** | **1024** | **569** | **657** | **7445** |

## HGVSp Category Breakdown

| Study | both_empty | delins_normalized | dup_vs_ins | fuzzy_fs | maf_empty | maf_nonstandard | match | mismatch | no_cds_data | position_shift | splice_no_protein | splice_vs_predicted | splice_vs_syn | transcript_model_change | vep_empty |
|-------|------|------|------|------|------|------|------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 479 | 7 | 1 | 645 | 739 | 2620 | 109994 | 0 | 74 | 489 | 57 | 28 | 611 | 47 | 59 |
| brca_tcga_gdc | 336 | 19 | 4 | 1040 | 447 | 1567 | 84167 | 0 | 79 | 659 | 94 | 45 | 476 | 25 | 54 |
| chol_tcga_gdc | 24 | 2 | 0 | 77 | 23 | 85 | 3484 | 0 | 4 | 30 | 8 | 1 | 21 | 1 | 4 |
| coad_tcga_gdc | 2317 | 8 | 2 | 3024 | 1306 | 5448 | 229187 | 0 | 196 | 1240 | 43 | 34 | 1372 | 88 | 287 |
| gbm_tcga_gdc | 199 | 0 | 2 | 582 | 303 | 1117 | 51900 | 0 | 58 | 305 | 14 | 11 | 326 | 22 | 31 |
| luad_tcga_gdc | 728 | 10 | 0 | 1060 | 1073 | 4118 | 181724 | 0 | 223 | 688 | 53 | 21 | 1005 | 62 | 103 |
| skcm_tcga_gdc | 1414 | 17 | 0 | 493 | 2009 | 7930 | 336733 | 0 | 399 | 1338 | 52 | 21 | 2673 | 125 | 246 |
| **Total** | **5497** | **63** | **9** | **6921** | **5900** | **22885** | **997189** | **0** | **1033** | **4749** | **321** | **161** | **6484** | **370** | **784** |

## HGVSc Category Breakdown

| Study | both_empty | delins_normalized | dup_vs_ins | maf_empty | match | position_shift | transcript_model_change | vep_empty |
|-------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 10 | 81 | 25 | 892 | 114177 | 367 | 298 | 0 |
| brca_tcga_gdc | 8 | 106 | 34 | 538 | 87775 | 402 | 149 | 0 |
| chol_tcga_gdc | 0 | 7 | 3 | 31 | 3695 | 18 | 10 | 0 |
| coad_tcga_gdc | 25 | 90 | 175 | 1653 | 240971 | 1205 | 431 | 2 |
| gbm_tcga_gdc | 3 | 9 | 22 | 369 | 54126 | 244 | 97 | 0 |
| luad_tcga_gdc | 25 | 207 | 45 | 1329 | 188218 | 606 | 438 | 0 |
| skcm_tcga_gdc | 47 | 152 | 26 | 2525 | 348715 | 1352 | 630 | 3 |
| **Total** | **118** | **652** | **330** | **7337** | **1037677** | **4194** | **2053** | **5** |

## Cancer Gene Mismatches

No mismatches across all 1207 cancer genes tested.

## Performance

Cache load time: 43.358s

| Study | Variants | Sequential | Seq v/s | Parallel | Par v/s | Speedup |
|-------|----------|-----------|---------|----------|---------|--------|
| blca_tcga_gdc | 115850 | 30.532s | 3794 | 15.054s | 7696 | 2.03x |
| brca_tcga_gdc | 89012 | 25.308s | 3517 | 11.584s | 7684 | 2.18x |
| chol_tcga_gdc | 3764 | 1.376s | 2735 | 475ms | 7924 | 2.90x |
| coad_tcga_gdc | 244552 | 51.286s | 4768 | 25.716s | 9510 | 1.99x |
| gbm_tcga_gdc | 54870 | 15.048s | 3646 | 5.711s | 9608 | 2.63x |
| luad_tcga_gdc | 190868 | 43.721s | 4366 | 23.271s | 8202 | 1.88x |
| skcm_tcga_gdc | 353450 | 1m14.442s | 4748 | 40.017s | 8833 | 1.86x |
| **Total** | **1052366** | **4m1.712s** | **4354** | **2m1.828s** | **8638** | **1.98x** |
