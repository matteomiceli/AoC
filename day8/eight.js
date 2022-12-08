const { getData } = require("../helpers/helpers");

const matrix = getData("./data.txt").array.map((line) => line.split(""));
const matrixHeight = matrix.length;
const matrixWidth = matrix[0].length;
const visibilityMap = {};
const scenicMap = {};

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

function getScenic(trees, point) {
  let scenic = 0;
  for (let i = 0; i < trees.length; i++) {
    const tree = trees[i];
    if (tree >= point) {
      scenic += 1;
      break;
    }
    scenic += 1;
  }
  return scenic;
}

function horizontalVisibility(x, y) {
  let visible = false;
  const point = matrix[y][x];

  const hLeft = checkToEdge(matrix[y].slice(0, x), point);
  const hRight = checkToEdge(matrix[y].slice(x + 1, matrixWidth), point);
  const scenicLeft = getScenic(matrix[y].slice(0, x).reverse(), point);
  const scenicRight = getScenic(matrix[y].slice(x + 1, matrixWidth), point);

  if (hRight || hLeft) {
    visible = true;
  }

  return { visible, scenic: { left: scenicLeft, right: scenicRight } };
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
  const scenicTop = getScenic(vertical.slice(0, y).reverse(), point);
  const scenicBot = getScenic(vertical.slice(y + 1, matrixHeight), point);

  if (vTop || vBot) {
    visible = true;
  }
  return { visible, scenic: { top: scenicTop, bot: scenicBot } };
}

matrix.forEach((line, y) => {
  visibilityMap[y] = {};
  scenicMap[y] = {};
  line.forEach((_, x) => {
    if (isEdge(x, y)) {
      visibilityMap[y][x] = true;
      return;
    }

    // who cares about code efficiency anyway 🙃
    const hzntl = horizontalVisibility(x, y);
    const vert = verticalVisibility(x, y);

    scenicMap[y][x] =
      hzntl.scenic.left *
      hzntl.scenic.right *
      vert.scenic.top *
      vert.scenic.bot;

    if (hzntl.visible) {
      visibilityMap[y][x] = true;
      return;
    }
    if (vert.visible) {
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

function getBestScenicScore(scenicMap) {
  let bestScore = 0;

  Object.values(scenicMap)
    .map((row) => Object.values(row))
    .forEach((row) => {
      row.forEach((score) => {
        if (score > bestScore) {
          bestScore = score;
        }
      });
    });

  return bestScore;
}

console.log(getBestScenicScore(scenicMap));

// console.log(visibilityMap);
// console.log(getTotalVisible(visibilityMap));
