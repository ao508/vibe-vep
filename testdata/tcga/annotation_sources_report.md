# Annotation Sources Report

Generated: 2026-03-01 04:57 UTC  
GENCODE transcripts: 254070  
AlphaMissense variants in database: 71697556  
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

## AlphaMissense Performance

| Study | Variants | Base Time | AM Lookup Time | AM Overhead | Lookups/sec |
|-------|----------|-----------|----------------|-------------|-------------|
| blca_tcga_gdc | 115850 | 19.151s | 5.03s | 26.3% | 15098 |
| brca_tcga_gdc | 89012 | 15.912s | 4.447s | 27.9% | 12977 |
| chol_tcga_gdc | 3764 | 679ms | 2.303s | 339.2% | 1015 |
| coad_tcga_gdc | 244552 | 30.406s | 7.233s | 23.8% | 17787 |
| gbm_tcga_gdc | 54870 | 8.464s | 3.649s | 43.1% | 8957 |
| luad_tcga_gdc | 190868 | 28.32s | 6.491s | 22.9% | 18526 |
| skcm_tcga_gdc | 353450 | 44.148s | 9.31s | 21.1% | 21953 |
| **Total** | **1052366** | **2m27.08s** | **38.461s** | **26.2%** | **16170** |
