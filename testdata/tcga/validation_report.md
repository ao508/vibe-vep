# TCGA Validation Report

Generated: 2026-03-05 01:31 UTC  
GENCODE transcripts loaded: 254070  
Cache load time: 32.727s  
Workers: 4 (GOMAXPROCS)

## Match Rates

| Study | Variants | Conseq Match | Conseq Mismatch | Conseq Rate | HGVSp Match | HGVSp Mismatch | HGVSp Rate | HGVSc Match | HGVSc Mismatch | HGVSc Rate |
|-------|----------|-------------|-----------------|-------------|-------------|----------------|------------|-------------|----------------|------------|
| blca_tcga_gdc | 115850 | 115647 | 1 | 99.8% | 109979 | 0 | 94.9% | 114159 | 0 | 98.5% |
| brca_tcga_gdc | 89012 | 88847 | 1 | 99.8% | 84151 | 0 | 94.5% | 87759 | 0 | 98.6% |
| chol_tcga_gdc | 3764 | 3757 | 0 | 99.8% | 3482 | 0 | 92.5% | 3693 | 0 | 98.1% |
| coad_tcga_gdc | 244552 | 244090 | 2 | 99.8% | 229154 | 0 | 93.7% | 240937 | 0 | 98.5% |
| gbm_tcga_gdc | 54870 | 54748 | 0 | 99.8% | 51901 | 0 | 94.6% | 54127 | 0 | 98.6% |
| luad_tcga_gdc | 190868 | 190446 | 1 | 99.8% | 181693 | 0 | 95.2% | 188186 | 0 | 98.6% |
| skcm_tcga_gdc | 353450 | 352555 | 3 | 99.7% | 336719 | 0 | 95.3% | 348702 | 0 | 98.7% |
| **Total** | **1052366** | **1050090** | **8** | **99.8%** | **997079** | **0** | **94.7%** | **1037563** | **0** | **98.6%** |

## Consequence Category Breakdown

| Study | delins_normalized | gene_model_change | match | mismatch | no_cds_data | position_shift | transcript_model_change | upstream_reclassified |
|-------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 1 | 0 | 114748 | 1 | 74 | 42 | 85 | 899 |
| brca_tcga_gdc | 2 | 0 | 88303 | 1 | 78 | 32 | 52 | 544 |
| chol_tcga_gdc | 0 | 0 | 3726 | 0 | 4 | 1 | 2 | 31 |
| coad_tcga_gdc | 8 | 2 | 242409 | 2 | 193 | 99 | 158 | 1681 |
| gbm_tcga_gdc | 0 | 0 | 54377 | 0 | 58 | 33 | 31 | 371 |
| luad_tcga_gdc | 1 | 0 | 189095 | 1 | 219 | 96 | 105 | 1351 |
| skcm_tcga_gdc | 5 | 3 | 349987 | 3 | 398 | 266 | 220 | 2568 |
| **Total** | **17** | **5** | **1042645** | **8** | **1024** | **569** | **653** | **7445** |

## HGVSp Category Breakdown

| Study | both_empty | delins_normalized | dup_vs_ins | fuzzy_fs | maf_empty | maf_nonstandard | match | mismatch | no_cds_data | position_shift | splice_no_protein | splice_vs_predicted | splice_vs_syn | transcript_model_change | vep_empty |
|-------|------|------|------|------|------|------|------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 479 | 7 | 1 | 645 | 739 | 2620 | 109979 | 0 | 74 | 504 | 57 | 28 | 611 | 47 | 59 |
| brca_tcga_gdc | 336 | 19 | 4 | 1040 | 447 | 1567 | 84151 | 0 | 79 | 675 | 94 | 45 | 476 | 25 | 54 |
| chol_tcga_gdc | 24 | 2 | 0 | 77 | 23 | 85 | 3482 | 0 | 4 | 32 | 8 | 1 | 21 | 1 | 4 |
| coad_tcga_gdc | 2317 | 8 | 2 | 3024 | 1306 | 5448 | 229154 | 0 | 196 | 1273 | 43 | 34 | 1372 | 88 | 287 |
| gbm_tcga_gdc | 199 | 0 | 2 | 582 | 303 | 1117 | 51901 | 0 | 58 | 304 | 14 | 11 | 326 | 22 | 31 |
| luad_tcga_gdc | 728 | 10 | 0 | 1060 | 1073 | 4118 | 181693 | 0 | 223 | 719 | 53 | 21 | 1005 | 62 | 103 |
| skcm_tcga_gdc | 1414 | 17 | 0 | 493 | 2009 | 7930 | 336719 | 0 | 399 | 1352 | 52 | 21 | 2673 | 125 | 246 |
| **Total** | **5497** | **63** | **9** | **6921** | **5900** | **22885** | **997079** | **0** | **1033** | **4859** | **321** | **161** | **6484** | **370** | **784** |

## HGVSc Category Breakdown

| Study | both_empty | delins_normalized | dup_vs_ins | maf_empty | match | position_shift | transcript_model_change | vep_empty |
|-------|------|------|------|------|------|------|------|------|
| blca_tcga_gdc | 10 | 81 | 25 | 892 | 114159 | 382 | 301 | 0 |
| brca_tcga_gdc | 8 | 106 | 34 | 538 | 87759 | 417 | 150 | 0 |
| chol_tcga_gdc | 0 | 7 | 3 | 31 | 3693 | 20 | 10 | 0 |
| coad_tcga_gdc | 25 | 90 | 175 | 1653 | 240937 | 1239 | 431 | 2 |
| gbm_tcga_gdc | 3 | 9 | 22 | 369 | 54127 | 243 | 97 | 0 |
| luad_tcga_gdc | 25 | 207 | 45 | 1329 | 188186 | 637 | 439 | 0 |
| skcm_tcga_gdc | 47 | 152 | 26 | 2525 | 348702 | 1364 | 631 | 3 |
| **Total** | **118** | **652** | **330** | **7337** | **1037563** | **4302** | **2059** | **5** |

## Cancer Gene Mismatches

No mismatches across all 1207 cancer genes tested.

## Performance

Cache load time: 32.727s

| Study | Variants | Sequential | Seq v/s | Parallel | Par v/s | Speedup |
|-------|----------|-----------|---------|----------|---------|--------|
| blca_tcga_gdc | 115850 | 23.122s | 5010 | 9.445s | 12265 | 2.45x |
| brca_tcga_gdc | 89012 | 21.714s | 4099 | 8.765s | 10156 | 2.48x |
| chol_tcga_gdc | 3764 | 937ms | 4016 | 412ms | 9147 | 2.28x |
| coad_tcga_gdc | 244552 | 41.13s | 5946 | 16.598s | 14734 | 2.48x |
| gbm_tcga_gdc | 54870 | 11.937s | 4597 | 4.553s | 12051 | 2.62x |
| luad_tcga_gdc | 190868 | 36.268s | 5263 | 14.798s | 12898 | 2.45x |
| skcm_tcga_gdc | 353450 | 58.912s | 6000 | 25.308s | 13966 | 2.33x |
| **Total** | **1052366** | **3m14.021s** | **5424** | **1m19.878s** | **13175** | **2.43x** |
