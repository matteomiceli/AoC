const { getData } = require("../helpers/helpers");

const matrix = getData("./data.txt").array.map((line) => line.split(""));
const matrixHeight = matrix.length;
const matrixWidth = matrix[0].length;
const visibilityMap = {};

// checks edges
function isEdge(x, y) {
  const top = matrix?.[y - 1]?.[x];
  const right = matrix?.[y]?.[x + 1];
  const bottom = matrix?.[y + 1]?.[x];
  const left = matrix?.[y]?.[x - 1];

  return [top, right, bottom, left].some((direction) => !direction);
}

function checkToEdge(trees, point) {
  return trees.every((tree) => {
    return point > tree;
  });
}

function horizontalVisibility(x, y) {
  let visible = false;
  const point = matrix[y][x];

  const hLeft = checkToEdge(matrix[y].slice(0, x), point);
  const hRight = checkToEdge(matrix[y].slice(x + 1, matrixWidth), point);

  if (hRight || hLeft) {
    visible = true;
  }

  return visible;
}

function verticalVisibility(x, y) {
  let visible = false;
  const point = matrix[y][x];

  const vertical = [];

  for (let i = 0; i < matrixHeight; i++) {
    vertical.push(matrix[i][x]);
  }

  const vTop = checkToEdge(vertical.slice(0, y), point);
  const vBot = checkToEdge(vertical.slice(y + 1, matrixHeight), point);

  if (vTop || vBot) {
    visible = true;
  }
  return visible;
}

matrix.forEach((line, y) => {
  visibilityMap[y] = {};
  line.forEach((_, x) => {
    if (isEdge(x, y)) {
      visibilityMap[y][x] = true;
      return;
    }
    if (horizontalVisibility(x, y)) {
      visibilityMap[y][x] = true;
      return;
    }
    if (verticalVisibility(x, y)) {
      visibilityMap[y][x] = true;
      return;
    }
    visibilityMap[y][x] = false;
  });
});

function getTotalVisible(visibilityMap) {
  let sum = 0;
  const visibilityArr = Object.values(visibilityMap).map((row) =>
    Object.values(row)
  );
  visibilityArr.forEach((row) => {
    row.forEach((point) => {
      if (point) sum += 1;
    });
  });
  return sum;
}

console.log(visibilityMap);
console.log(getTotalVisible(visibilityMap));
