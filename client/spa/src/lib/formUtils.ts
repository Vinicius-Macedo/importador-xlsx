export interface FormValue {
  [key: string]: string;
}

export function handleFormsInputChange(
  event: React.ChangeEvent<HTMLInputElement>,
  form: FormValue,
  setForm: (React.Dispatch<React.SetStateAction<FormValue>>)
) {
  setForm({
    ...form,
    [event.target.name]: event.target.value,
  });
}
