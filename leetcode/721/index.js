let accounts = [
  ["John", "johnsmith@mail.com", "john00@mail.com"],
  ["John", "johnnybravo@mail.com"],
  ["Johnson", "johnsmith@mail.com", "john_newyork@mail.com"],
  ["Mary", "mary@mail.com"]
];

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

let text = [
  [0, 3],
  [1, 2],
  [3, 5]
];

uf = new UnionFind(6);

console.log(uf.count, uf.parent, uf.rank);

for (let pair of text) {
  uf.union(pair[0], pair[1]);
}

console.log(uf.count, uf.parent, uf.rank);

/**
 * @param {string[][]} accounts
 * @return {string[][]}
 */
var accountsMerge = function (accounts) {
  let res = [];
  let idToName = new Map();
  let emailToID = new Map();
  let idToEmails = new Map();

  uf = new UnionFind(accounts.length);

  for (let id = 0; id < accounts.length; id++) {
    idToName.set(id, accounts[id][0]);
    for (let i = 1; i < accounts[id].length; i++) {
      let pid = emailToID.get(accounts[id][i]);

      if (pid !== undefined) {
        uf.union(pid, id);
      }

      emailToID.set(accounts[id][i], id);
    }
  }

  for (let [email, id] of emailToID) {
    const rootID = uf.find(id);
    let emails = idToEmails.get(rootID);
    if (emails) {
      emails.push(email);
      idToEmails.set(rootID, emails);
    } else {
      idToEmails.set(rootID, [email]);
    }
  }

  for (let [id, emails] of idToEmails) {
    emails = emails.sort();
    const account = [idToName.get(id), ...emails];
    res.push(account);
  }

  return res;
};

console.log(accountsMerge(accounts));
