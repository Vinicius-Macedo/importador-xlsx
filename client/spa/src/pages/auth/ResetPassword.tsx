import { useLocation } from "react-router-dom";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState, useEffect } from "react";
import { handleFormsInputChange, FormValue } from "@/lib/formUtils";
import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { useNavigate } from "react-router-dom";
import { validateFields } from "@/helpers/formValidatorHelper";
import PasswordHelper from "@/components/passwordHelper";

const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};

export function ResetPassword() {
  const query = useQuery();

  const [form, setForm] = useState<FormValue>({
    token: "",
    password: "",
    repeat_password: "",
  });
  const [error, setError] = useState<string>("");
  const [fieldsError, setFieldsError] = useState<{
    [key: string]: string;
  }>({});
  const [isPasswordFocused, setIsPasswordFocused] = useState<boolean>(false);

  useEffect(() => {
    const token = query.get("token");
    if (token) {
      setForm({ ...form, token });
    }
  }, []);

  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: () =>
      axiosInstance.post("/recover-password", {
        token: form.token,
        password: form.password,
      }),
    onSuccess: () => {
      navigate("/success-password");
    },
    onError: (error: any) => {
      setError(
        error.response?.data?.error ||
          "Ocorreu um erro, por favor entre em contato com o suporte"
      );
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (Object.keys(validateFields(form)).length > 0) {
      console.log("validateFields", validateFields(form));
      setError("Corrija os campos inválidos");
      setFieldsError(validateFields(form));
      return;
    }

    mutation.mutate();
  };

  return (
    <>
      <div className="flex justify-center items-center">
        <Card className="w-full max-w-sm">
          <form onSubmit={(e) => handleSubmit(e)}>
            <CardHeader>
              <CardTitle className="text-2xl text-center">
                Esqueceu sua senha!
              </CardTitle>
              <CardDescription className="text-center">
                Digite a nova senha
                {error && <p className="text-red-500">{error}</p>}
              </CardDescription>
            </CardHeader>
            <CardContent className="grid gap-4">
              <div className="grid gap-2 relative">
                <Label htmlFor="password">Nova senha</Label>
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
                  <p className="text-red-500">{fieldsError.password}</p>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="confirm-password">Confirme a senha</Label>
                <Input
                  id="confirm-password"
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
            </CardContent>
            <CardFooter className="flex flex-col items-center gap-4">
              <Button disabled={mutation.isPending}>
                {mutation.isPending ? "Carregando..." : "Enviar"}
              </Button>
            </CardFooter>
          </form>
        </Card>
      </div>
    </>
  );
}
