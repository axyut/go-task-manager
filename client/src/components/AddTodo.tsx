import { useState } from "react";
import { useForm } from "@mantine/form";
import { Button, Group, Modal, TextInput, Textarea } from "@mantine/core";
import { KeyedMutator } from "swr";
import { API_URL, Todo } from "../App";

export default function AddTodo({ mutate }: { mutate: KeyedMutator<Todo[]> }) {
  const [open, setOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  async function createTodo(values: { title: string; body: string }) {
    const updated = await fetch(`${API_URL}todos`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
    }).then((res) => res.json());
    mutate(updated);
    form.reset();
    setOpen(false);
  }
  return (
    <div>
      <Modal opened={open} onClose={() => setOpen(false)} title="Create Task">
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Task"
            placeholder="I need to?"
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Description"
            placeholder="Explain[10]"
            {...form.getInputProps("body")}
          />
          <Button type="submit">Create</Button>
        </form>
      </Modal>
      <Group position="center">
        {" "}
        <Button fullWidth mb={12} onClick={() => setOpen(true)}>
          Add Task
        </Button>
      </Group>
    </div>
  );
}
