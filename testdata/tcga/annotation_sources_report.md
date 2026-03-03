# Annotation Sources Report

Generated: 2026-03-03 17:26 UTC  
GENCODE transcripts: 254070  
AlphaMissense variants in database: 71697556  
Cancer Hotspots: 682 transcripts, 4189 positions  
ClinVar variants: 4144490  
SIGNAL variants: 92997 (GRCh37 coordinates)  
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
| blca_tcga_gdc | 115850 | 114037 | 162 | 0.14% | 99 | 0 | 54 | 9 |
| brca_tcga_gdc | 89012 | 87866 | 451 | 0.51% | 358 | 0 | 89 | 4 |
| chol_tcga_gdc | 3764 | 3695 | 12 | 0.32% | 10 | 0 | 2 | 0 |
| coad_tcga_gdc | 244552 | 239629 | 379 | 0.16% | 181 | 0 | 196 | 2 |
| gbm_tcga_gdc | 54870 | 54144 | 108 | 0.20% | 39 | 0 | 61 | 8 |
| luad_tcga_gdc | 190868 | 187939 | 220 | 0.12% | 62 | 0 | 145 | 13 |
| skcm_tcga_gdc | 353450 | 348150 | 423 | 0.12% | 223 | 0 | 198 | 2 |
| **Total** | **1052366** | **1035460** | **1755** | **0.17%** | **972** | **0** | **745** | **38** |

## ClinVar Coverage

| Study | Variants | ClinVar Hits | Hit Rate | Pathogenic | Likely_pathogenic | Uncertain | Benign | Likely_benign | Other |
|-------|----------|--------------|----------|------------|-------------------|-----------|--------|---------------|-------|
| blca_tcga_gdc | 115850 | 8135 | 7.02% | 562 | 144 | 4985 | 47 | 1648 | 749 |
| brca_tcga_gdc | 89012 | 7709 | 8.66% | 730 | 139 | 4515 | 55 | 1478 | 792 |
| chol_tcga_gdc | 3764 | 366 | 9.72% | 21 | 6 | 225 | 2 | 69 | 43 |
| coad_tcga_gdc | 244552 | 30838 | 12.61% | 1428 | 460 | 19163 | 213 | 6613 | 2961 |
| gbm_tcga_gdc | 54870 | 8035 | 14.64% | 419 | 102 | 4945 | 59 | 1673 | 837 |
| luad_tcga_gdc | 190868 | 9145 | 4.79% | 532 | 295 | 5607 | 54 | 1858 | 799 |
| skcm_tcga_gdc | 353450 | 28274 | 8.00% | 984 | 492 | 17033 | 148 | 7318 | 2299 |
| **Total** | **1052366** | **92502** | **8.79%** | **4676** | **1638** | **56473** | **578** | **20657** | **8480** |

## SIGNAL Coverage

*Note: SIGNAL data uses GRCh37 coordinates. TCGA GDC MAFs use GRCh38, so few hits are expected.*

| Study | Variants | SIGNAL Hits | Hit Rate |
|-------|----------|-------------|----------|
| blca_tcga_gdc | 115850 | 9 | 0.01% |
| brca_tcga_gdc | 89012 | 2 | 0.00% |
| chol_tcga_gdc | 3764 | 0 | 0.00% |
| coad_tcga_gdc | 244552 | 27 | 0.01% |
| gbm_tcga_gdc | 54870 | 6 | 0.01% |
| luad_tcga_gdc | 190868 | 8 | 0.00% |
| skcm_tcga_gdc | 353450 | 40 | 0.01% |
| **Total** | **1052366** | **92** | **0.01%** |

## AlphaMissense Performance

| Study | Variants | Base Time | AM Lookup Time | AM Overhead | Lookups/sec |
|-------|----------|-----------|----------------|-------------|-------------|
| blca_tcga_gdc | 115850 | 22.866s | 12.675s | 55.4% | 5991 |
| brca_tcga_gdc | 89012 | 21.598s | 6.983s | 32.3% | 8263 |
| chol_tcga_gdc | 3764 | 657ms | 3.755s | 571.2% | 622 |
| coad_tcga_gdc | 244552 | 51.512s | 10.729s | 20.8% | 11991 |
| gbm_tcga_gdc | 54870 | 12.18s | 6.714s | 55.1% | 4868 |
| luad_tcga_gdc | 190868 | 42.129s | 10.189s | 24.2% | 11801 |
| skcm_tcga_gdc | 353450 | 55.635s | 12.461s | 22.4% | 16401 |
| **Total** | **1052366** | **3m26.577s** | **1m3.507s** | **30.7%** | **9793** |
