# Annotation Sources Report

Generated: 2026-03-11 14:49 UTC  
GENCODE transcripts: 254070  
Data: unified genomic index (SQLite)  
Workers: 4 (GOMAXPROCS)

## AlphaMissense Coverage

| Study | Variants | Missense | AM Hits | Coverage | likely_benign | ambiguous | likely_pathogenic |
|-------|----------|----------|---------|----------|---------------|-----------|-------------------|
| blca_tcga_gdc | 115850 | 75938 | 69937 | 92.1% | 41805 | 7662 | 20470 |
| brca_tcga_gdc | 89012 | 57703 | 53240 | 92.3% | 31416 | 5807 | 16017 |
| chol_tcga_gdc | 3764 | 2336 | 2151 | 92.1% | 1244 | 234 | 673 |
| coad_tcga_gdc | 244552 | 128655 | 118282 | 91.9% | 70342 | 12727 | 35213 |
| gbm_tcga_gdc | 54870 | 32684 | 30155 | 92.3% | 18286 | 3259 | 8610 |
| luad_tcga_gdc | 190868 | 120249 | 111597 | 92.8% | 65087 | 12682 | 33828 |
| skcm_tcga_gdc | 353450 | 204372 | 185786 | 90.9% | 115483 | 20014 | 50289 |
| **Total** | **1052366** | **621937** | **571148** | **91.8%** | **343663** | **62385** | **165100** |

## ClinVar Coverage

| Study | Variants | ClinVar Hits | Hit Rate | Pathogenic | Likely_pathogenic | Uncertain | Benign | Likely_benign | Other |
|-------|----------|--------------|----------|------------|-------------------|-----------|--------|---------------|-------|
| blca_tcga_gdc | 115850 | 8211 | 7.09% | 605 | 153 | 4999 | 47 | 1649 | 758 |
| brca_tcga_gdc | 89012 | 8025 | 9.02% | 915 | 167 | 4559 | 56 | 1478 | 850 |
| chol_tcga_gdc | 3764 | 395 | 10.49% | 39 | 9 | 230 | 2 | 70 | 45 |
| coad_tcga_gdc | 244552 | 32906 | 13.46% | 2415 | 776 | 19501 | 224 | 6643 | 3347 |
| gbm_tcga_gdc | 54870 | 8298 | 15.12% | 549 | 121 | 4984 | 59 | 1680 | 905 |
| luad_tcga_gdc | 190868 | 9334 | 4.89% | 622 | 311 | 5635 | 54 | 1859 | 853 |
| skcm_tcga_gdc | 353450 | 28341 | 8.02% | 1023 | 498 | 17041 | 148 | 7319 | 2312 |
| **Total** | **1052366** | **95510** | **9.08%** | **6168** | **2035** | **56949** | **590** | **20698** | **9070** |

## SIGNAL Coverage

*Note: SIGNAL data uses GRCh37 coordinates. TCGA GDC MAFs use GRCh38, so few hits are expected.*

| Study | Variants | SIGNAL Hits | Hit Rate |
|-------|----------|-------------|----------|
| blca_tcga_gdc | 115850 | 0 | 0.00% |
| brca_tcga_gdc | 89012 | 0 | 0.00% |
| chol_tcga_gdc | 3764 | 0 | 0.00% |
| coad_tcga_gdc | 244552 | 0 | 0.00% |
| gbm_tcga_gdc | 54870 | 0 | 0.00% |
| luad_tcga_gdc | 190868 | 0 | 0.00% |
| skcm_tcga_gdc | 353450 | 0 | 0.00% |
| **Total** | **1052366** | **0** | **0.00%** |

## Genomic Index Lookup Performance

| Study | Variants | Base Time | Index Lookup Time | Overhead | Lookups/sec |
|-------|----------|-----------|-------------------|----------|-------------|
| blca_tcga_gdc | 115850 | 33.494s | 9m13.761s | 1653.3% | 209 |
| brca_tcga_gdc | 89012 | 24.691s | 3m11.473s | 775.5% | 465 |
| chol_tcga_gdc | 3764 | 1.171s | 7.803s | 666.6% | 482 |
| coad_tcga_gdc | 244552 | 44.747s | 4m40.477s | 626.8% | 872 |
| gbm_tcga_gdc | 54870 | 10.466s | 37.869s | 361.8% | 1449 |
| luad_tcga_gdc | 190868 | 28.542s | 1m51.628s | 391.1% | 1710 |
| skcm_tcga_gdc | 353450 | 43.271s | 1m34.7s | 218.9% | 3732 |
| **Total** | **1052366** | **3m6.382s** | **21m17.711s** | **685.5%** | **824** |
