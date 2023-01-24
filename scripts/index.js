function add(x, y, p) {
  return (x + y) % p;
}

function sub(x, y, p) {
  let z = (x - y) % p;
  if (z < 0n) {
      z += p;
  }
  return z;
}

function mul(x, y, p) {
  return (x * y) % p;
}

function inv(x, p) {
  return pow(x, p - 2n, p);
}

function pow(x, y, p) {
  let z = 1n;
  while (y > 0n) {
      if (y & 1n) {
          z = z * x % p;
      }
      x = (x * x) % p;
      y >>= 1n;
  }
  return z;
}

function addPoint(P1, P2) {
  if (P1.x == null && P1.y == null) {
    return P2;
  }
  if (P2.x == null && P2.y == null) {
    return P1;
  }

  let s = null;
  if (P1.x != P2.x) {
    // 異なる点
    s = sub(P2.y, P1.y, p) * inv(sub(P2.x, P1.x, p), p);
  } else {
    // 同一点
    if (P1.y != P2.y) {
      // P1 = -P2
      return { x: null, y: null };
    } else if ((P1.y == 0n)) {
      // 接線が垂直
      return { x: null, y: null };
    }
    s = mul(3n, pow(P1.x, 2n, p), p) * inv(mul(2n, P1.y, p), p);
  }

  let x3 = sub(sub(pow(s, 2n, p), P1.x, p), P2.x, p);
  let y3 = sub(mul(s, sub(P1.x, x3, p), p), P1.y, p);

  return { x: x3, y: y3 };
}

function mulPoint(P, k) {
  let Q = { x: null, y: null };

  while(k > 0n) {
    if (k & 1n) {
      Q = addPoint(Q, P);
    }
    P = addPoint(P, P);
    k >>= 1n;
  }
  return Q;
}

async function main() {
  // 楕円曲線係数（ secp256k1 : y^2 = x^3 + 7　）
  const b = 7n;

  // 楕円曲線の位数
  const p = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F');

  // 生成元
  const G = {
    x: BigInt('0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798'),
    y: BigInt('0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8'),
  }

  // 生成元の位数
  const n = BigInt('0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141');
}

main()