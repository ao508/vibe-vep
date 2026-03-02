# TCGA Validation Report

Generated: 2026-03-02 02:38 UTC  
GENCODE transcripts loaded: 254070  
Cache load time: 29.293s  
Workers: 4 (GOMAXPROCS)

## Match Rates

| Study | Variants | Conseq Match | Conseq Mismatch | Conseq Rate | HGVSp Match | HGVSp Mismatch | HGVSp Rate | HGVSc Match | HGVSc Mismatch | HGVSc Rate |
|-------|----------|-------------|-----------------|-------------|-------------|----------------|------------|-------------|----------------|------------|
| chol_tcga_gdc | 3764 | 3757 | 0 | 99.8% | 3483 | 0 | 92.5% | 3694 | 0 | 98.1% |
| **Total** | **3764** | **3757** | **0** | **99.8%** | **3483** | **0** | **92.5%** | **3694** | **0** | **98.1%** |

## Consequence Category Breakdown

| Study | match | mismatch | no_cds_data | position_shift | transcript_model_change | upstream_reclassified |
|-------|------|------|------|------|------|------|
| chol_tcga_gdc | 3726 | 0 | 4 | 1 | 2 | 31 |
| **Total** | **3726** | **0** | **4** | **1** | **2** | **31** |

## HGVSp Category Breakdown

| Study | both_empty | delins_normalized | fuzzy_fs | maf_empty | maf_nonstandard | match | mismatch | no_cds_data | position_shift | splice_no_protein | splice_vs_predicted | splice_vs_syn | transcript_model_change | vep_empty |
|-------|------|------|------|------|------|------|------|------|------|------|------|------|------|------|
| chol_tcga_gdc | 24 | 2 | 77 | 23 | 85 | 3483 | 0 | 4 | 31 | 8 | 1 | 21 | 1 | 4 |
| **Total** | **24** | **2** | **77** | **23** | **85** | **3483** | **0** | **4** | **31** | **8** | **1** | **21** | **1** | **4** |

## HGVSc Category Breakdown

| Study | delins_normalized | dup_vs_ins | maf_empty | match | position_shift | transcript_model_change |
|-------|------|------|------|------|------|------|
| chol_tcga_gdc | 7 | 3 | 31 | 3694 | 19 | 10 |
| **Total** | **7** | **3** | **31** | **3694** | **19** | **10** |

## Cancer Gene Mismatches

No mismatches across all 312 cancer genes tested.

## Performance

Cache load time: 29.293s

| Study | Variants | Sequential | Seq v/s | Parallel | Par v/s | Speedup |
|-------|----------|-----------|---------|----------|---------|--------|
| chol_tcga_gdc | 3764 | 826ms | 4559 | 303ms | 12425 | 2.73x |
| **Total** | **3764** | **826ms** | **4559** | **303ms** | **12425** | **2.73x** |
