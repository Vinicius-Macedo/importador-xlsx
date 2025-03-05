import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { handleFormsInputChange, FormValue } from "@/lib/formUtils";
import { useState } from "react";
import { axiosInstance } from "@/services/axiosInstance";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import PasswordHelper from "@/components/passwordHelper";
import { validateFields } from "@/helpers/formValidatorHelper";
import { Check } from "lucide-react";
import { toast } from "sonner";

export function Register() {
  const [form, setForm] = useState<FormValue>({
    name: "",
    username: "",
    email: "",
    password: "",
    repeat_password: "",
  });
  const [error, setError] = useState<string>("");
  const [fieldsError, setFieldsError] = useState<{
    [key: string]: string;
  }>({});
  const [isPasswordFocused, setIsPasswordFocused] = useState<boolean>(false);

  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: () =>
      axiosInstance.post("/register", {
        name: form.name,
        username: form.username,
        email: form.email,
        password: form.password,
      }),
    onSuccess: () => {
      toast("Conta criada com sucesso!", {
        icon: <Check />,
      });
      navigate("/login");
    },
    onError: (error: any) => {
      setError(error.response?.data?.error || "An error occurred");
    },
  });

  function handleSubmit(event: React.FormEvent) {
    event.preventDefault();

    if (Object.keys(validateFields(form)).length > 0) {
      setError("Corrija os campos inválidos");
      setFieldsError(validateFields(form));
      return;
    }

    if (form.password !== form.repeat_password) {
      alert("Passwords do not match");
      return;
    }

    mutation.mutate();
  }

  return (
    <div className="flex flex-1 justify-center items-center">
      <div className="p-8 md:w-[400px] flex flex-col gap-4">
        <div>
          <p className="text-3xl font-bold text-center">Criar Conta</p>
          <p className="text-balance text-muted-foreground text-center">
            {error ? (
              <span className="text-red-500">{error}</span>
            ) : (
              "Preencha os campos para criar uma conta"
            )}
          </p>
        </div>
        <div>
          <form onSubmit={(e) => handleSubmit(e)}>
            <div className="grid gap-4">
              <div className="grid gap-2">
                <Label htmlFor="name">Seu Nome</Label>
                <Input
                  id="name"
                  type="text"
                  name="name"
                  onChange={(e) =>
                    handleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
                {fieldsError.name && (
                  <p className="text-red-500 text-sm">{fieldsError.name}</p>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="username">Nome de Usuário</Label>
                <Input
                  id="username"
                  type="text"
                  name="username"
                  onChange={(e) =>
                    handleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
                {fieldsError.username && (
                  <p className="text-red-500 text-sm">{fieldsError.username}</p>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  name="email"
                  onChange={(e) =>
                    handleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
                {fieldsError.email && (
                  <p className="text-red-500 text-sm">{fieldsError.email}</p>
                )}
              </div>
              <div className="grid gap-2 relative">
                <Label htmlFor="password">Senha</Label>
                <Input
                  id="password"
                  type="password"
                  name="password"
                  onChange={(e) =>
                    handleFormsInputChange(e, form as FormValue, setForm)
                  }
                  onFocus={() => setIsPasswordFocused(true)}
                  onBlur={() => setIsPasswordFocused(false)}
                />
                <PasswordHelper
                  password={form.password}
                  focus={isPasswordFocused}
                />
                {fieldsError.password && (
                  <p className="text-red-500 text-sm">{fieldsError.password}</p>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="repeat-password">Repita a Senha</Label>
                <Input
                  id="repeat-password"
                  type="password"
                  name="repeat_password"
                  onChange={(e) =>
                    handleFormsInputChange(e, form as FormValue, setForm)
                  }
                />
                {fieldsError.repeat_password && (
                  <p className="text-red-500">{fieldsError.repeat_password}</p>
                )}
              </div>
              <Button
                disabled={mutation.isPending}
                type="submit"
                className="w-full"
              >
                {mutation.isPending ? "Carregando..." : "Criar Conta"}
              </Button>
            </div>
            <p className="mt-4 text-center text-sm">
              Já tem uma conta?{" "}
              <span
                onClick={() => navigate("/login")}
                className="underline cursor-pointer"
              >
                Entrar
              </span>
            </p>
          </form>
        </div>
      </div>
    </div>
  );
}
