# Annotation Sources Report

Generated: 2026-03-05 01:21 UTC  
GENCODE transcripts: 254070  
Data: unified genomic index (SQLite)  
Cancer Hotspots: 682 transcripts, 4189 positions  
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

## Cancer Hotspots Coverage

| Study | Variants | Checked | Hotspot Hits | Hit Rate | single residue | in-frame indel | 3d | splice site |
|-------|----------|---------|--------------|----------|----------------|----------------|----|-------------|
| blca_tcga_gdc | 115850 | 114037 | 140 | 0.12% | 67 | 0 | 73 | 0 |
| brca_tcga_gdc | 89012 | 87866 | 266 | 0.30% | 182 | 0 | 79 | 5 |
| chol_tcga_gdc | 3764 | 3695 | 8 | 0.22% | 7 | 0 | 1 | 0 |
| coad_tcga_gdc | 244552 | 239629 | 301 | 0.13% | 111 | 0 | 189 | 1 |
| gbm_tcga_gdc | 54870 | 54144 | 47 | 0.09% | 15 | 0 | 31 | 1 |
| luad_tcga_gdc | 190868 | 187939 | 194 | 0.10% | 47 | 0 | 138 | 9 |
| skcm_tcga_gdc | 353450 | 348150 | 394 | 0.11% | 198 | 0 | 194 | 2 |
| **Total** | **1052366** | **1035460** | **1350** | **0.13%** | **627** | **0** | **705** | **18** |

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
| blca_tcga_gdc | 115850 | 24.147s | 56.354s | 233.4% | 2056 |
| brca_tcga_gdc | 89012 | 21.406s | 39.398s | 184.1% | 2259 |
| chol_tcga_gdc | 3764 | 953ms | 1.932s | 202.8% | 1948 |
| coad_tcga_gdc | 244552 | 41.399s | 1m39.291s | 239.8% | 2463 |
| gbm_tcga_gdc | 54870 | 11.816s | 25.648s | 217.1% | 2139 |
| luad_tcga_gdc | 190868 | 34.229s | 1m12.945s | 213.1% | 2617 |
| skcm_tcga_gdc | 353450 | 57.599s | 2m14.805s | 234.0% | 2622 |
| **Total** | **1052366** | **3m11.548s** | **7m10.374s** | **224.7%** | **2445** |
