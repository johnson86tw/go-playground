class UnionFind {
  constructor(n) {
    this.count = n;
    this.parent = [];
    this.rank = [];

    for (let i = 0; i < n; i++) {
      this.parent[i] = i;
      this.rank[i] = 0;
    }
  }

  find(p) {
    let root = p;

    while (root !== this.parent[root]) {
      root = this.parent[root];
    }

    // path compression
    while (this.parent[p] !== root) {
      this.parent[p] = root;
    }

    return root;
  }

  union(p, q) {
    let rootp = this.find(p);
    let rootq = this.find(q);

    if (this.rank[rootp] < this.rank[rootq]) {
      this.parent[rootp] = rootq;
    } else {
      this.parent[rootq] = rootp;
      if (this.rank[rootp] === this.rank[rootq]) {
        this.rank[rootp]++;
      }
    }

    this.count--;
  }
}
