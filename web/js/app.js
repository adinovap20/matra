const btnAnalyze = document.getElementById("analyze-button");

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

loadMatra();
