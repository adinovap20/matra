const btnAnalyze = document.getElementById("analyze-button");
const txtSanskrit = document.getElementById("sanskrit-text");
const secResult = document.getElementById("results");

async function loadMatra() {
  try {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
      fetch("wasm/matra.wasm"),
      go.importObject,
    );

    go.run(result.instance);

    btnAnalyze.textContent = "Analyze";
    btnAnalyze.disabled = false;
  } catch (error) {
    console.error(error);
  }
}

function renderResults(lines) {
  secResult.replaceChildren();

  const copyAksharasBtn = document.createElement("button");
  copyAksharasBtn.textContent = "Copy Aksharas";
  copyAksharasBtn.addEventListener("click", async () => {
    const text = lines
      .filter((line) => line.text !== "")
      .map((line) => line.aksharas.join("\t"))
      .join("\n");

    try {
      await navigator.clipboard.writeText(text);
      copyAksharasBtn.textContent = "Copied!";
    } catch (error) {
      console.error("Failed to copy aksharas:", error);
      copyAksharasBtn.textContent = "Copy failed";
    }
  });

  const gap = document.createElement("p");
  secResult.append(copyAksharasBtn, gap);

  lines.forEach((line, index) => {
    if (line.text === "") {
      return;
    }

    const article = document.createElement("article");

    const heading = document.createElement("h3");
    heading.textContent = `Line ${index + 1}`;

    const text = document.createElement("p");
    text.textContent = line.line;

    const aksharaCount = document.createElement("p");
    aksharaCount.textContent = `Aksharas Count: ${line.aksharasCount}`;

    const matraCount = document.createElement("p");
    matraCount.textContent = `Matra Count: ${line.matraCount}`;

    const table = document.createElement("table");
    const tbody = document.createElement("tbody");

    // Akshara row
    const aksharaRow = document.createElement("tr");
    const aksharaLabel = document.createElement("th");
    aksharaLabel.textContent = "Akshara:";
    aksharaRow.appendChild(aksharaLabel);

    // Cumulative Akshara row
    const countRow = document.createElement("tr");
    const countLabel = document.createElement("th");
    countLabel.textContent = "Cumulative Akshara:";
    countRow.appendChild(countLabel);

    // Matra Row
    const matraRow = document.createElement("tr");
    const matraLabel = document.createElement("th");
    matraLabel.textContent = "Matra:";
    matraRow.appendChild(matraLabel);

    // Cumulative Matra Row
    const cumulativeMatraRow = document.createElement("tr");
    const cumulativeMatraLabel = document.createElement("th");
    cumulativeMatraLabel.textContent = "Cumulative Matra:";
    cumulativeMatraRow.appendChild(cumulativeMatraLabel);

    line.aksharas.forEach((akshara, i) => {
      const aksharaCell = document.createElement("td");
      aksharaCell.textContent = akshara;
      aksharaRow.appendChild(aksharaCell);

      const countCell = document.createElement("td");
      countCell.textContent = i + 1;
      countRow.appendChild(countCell);
    });

    let sum = 0;
    line.matras.forEach((matra, i) => {
      const matraCell = document.createElement("td");
      matraCell.textContent = matra;
      matraRow.appendChild(matraCell);

      sum += matra;

      const cumulativeMatraCell = document.createElement("td");
      cumulativeMatraCell.textContent = sum;
      cumulativeMatraRow.append(cumulativeMatraCell);
    });

    tbody.append(aksharaRow, countRow, matraRow, cumulativeMatraRow);
    table.appendChild(tbody);

    const tableContainer = document.createElement("div");
    tableContainer.style.overflowX = "auto";
    tableContainer.style.maxWidth = "100%";

    tableContainer.appendChild(table);
    article.append(heading, text, aksharaCount, matraCount, tableContainer);
    secResult.appendChild(article);
  });
}

function analyze() {
  const text = txtSanskrit.value;
  if (!text) {
    alert("Please enter or paste a sanskrit verse.");
    return;
  }

  secResult.hidden = false;

  try {
    const resultJson = window.analyzeSanskrit(text);
    const result = JSON.parse(resultJson);
    renderResults(result.analysis);
  } catch (error) {
    console.log(error);
  }
}

loadMatra();

btnAnalyze.addEventListener("click", () => {
  analyze();
});
