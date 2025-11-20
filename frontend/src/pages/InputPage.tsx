import { useState } from 'react';
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"

interface InputPageProps {
  onSubmit: (value: string) => void;
}

function InputPage({ onSubmit } : InputPageProps) {
  const [input, setInput] = useState<string>("");

  const handleSubmit = () => {
    if (input.trim() !== "") {
      onSubmit(input);
    }
  };

  return(
    <div className="grid grid-cols-1 place-items-center justify-items-center mx-auto py-8 space-y-6">
      <Textarea
        className="resize-none w-4/5 md:text-2xl h-[256px]"
        value={input}
        onChange={e => setInput(e.target.value)}
        placeholder="Введите сюда текст для анализа"
      />
      <Button
        onClick={handleSubmit}
        disabled={!input.trim()}
      >
        Обработать
      </Button>
    </div>
  );
}

export default InputPage;
