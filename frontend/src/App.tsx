import './App.css';
import InputPage from './pages/InputPage';
import ResultsPage from './pages/ResultsPage';
import { useState, useEffect } from 'react';
import { Parse } from '../wailsjs/go/app/App';
import { types } from '../wailsjs/go/models';

function App() {
  const [input, setInput] = useState<string>("");
  const [parserResult, setParserResult] = useState<types.ParserResult | null>(null);

  useEffect(() => {
    if (!input) {
      return;
    }
    const run = async () => {
      const res = await Parse(input);
      setParserResult(res);
    };
    run();
  }, [input]);

  if (input === "") {
    return (
      <InputPage onSubmit={(v: string) => setInput(v)} />
    );
  } else {
    return (
      <ResultsPage parserResult={parserResult} reset={() => setInput("")} />
    )
  }
}

export default App;
