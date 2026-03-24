# Event Stream DSA – Problem Roadmap

Dataset: Retail Rocket (≈250MB, ~2.7M events)  
Goal: Learn DSA (sorting, searching, stack, queue) via real-world event stream processing (no DB)

---

## 0. Streaming Engine
Process large dataset without loading into memory.

- Input: CSV (250MB)
- Output: streamed events
- DSA: stream processing, iterator pattern
- Key: O(1) memory

---

## 1. Most Viewed Products
Find most frequently viewed items.

- Filter: event = "view"
- DSA: HashMap (count), Sorting / Top-K

---

## 2. Most Purchased Products
Find most purchased items.

- Filter: event = "transaction"
- DSA: HashMap, Sorting

---

## 3. Conversion Rate per Product
Measure view → purchase ratio.

- Track: views & purchases per item
- DSA: HashMap, search
- Output: conversion rate

---

## 4. Abandoned Cart Detection
Detect add-to-cart without purchase.

- Flow: addtocart → (no transaction)
- DSA: HashMap (state tracking)

---

## 5. User Session Reconstruction
Rebuild user activity flow.

- Example: view → view → addtocart → purchase
- DSA: Stack (per user)

---

## 6. Most Common User Paths
Find most frequent navigation sequences.

- Example: A → B → C
- DSA: Stack + HashMap (path counting)

---

## 7. Trending Products (Sliding Window)
Detect products with sudden spike.

- Window: last N minutes
- DSA: Queue (sliding window)

---

## 8. Traffic Spike Detection
Detect high traffic periods.

- Example: > X events / minute
- DSA: Queue

---

## 9. Suspicious User Detection
Detect abnormal user activity.

- Example: 100 events in 5 sec
- DSA: Queue per user

---

## 10. Fast Event Lookup (Indexing)
Enable fast search by user/product.

- Without index: O(n)
- With index: O(1)
- DSA: HashMap (index)

---

## 11. External Sorting (Advanced)
Sort large dataset without full memory load.

- DSA: Merge sort (external)

---

## 12. Mini Analytics Engine (Final)
Combine all analyzers:

- most viewed
- trending
- sessions
- conversion

Goal:
Build a streaming analytics engine using basic data structures.

---

## DSA Mapping Summary

- HashMap → counting, indexing
- Stack → session / navigation flow
- Queue → sliding window, real-time detection
- Sorting → ranking, top-K
- Stream → memory efficiency

---

## Key Insight

Process:
read → parse → process → discard

Result:
- Dataset: 250MB
- Memory: ~3MB
- Throughput: ~700k+ events/sec

Streaming > loading everything into memory
