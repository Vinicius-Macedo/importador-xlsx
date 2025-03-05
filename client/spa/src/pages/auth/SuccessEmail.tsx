import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export function SuccessEmail() {
  const navigate = useNavigate();

  return (
    <>
      <section className="justify-center flex h-screen items-center">
        <div className="container">
          <div className='flex items-center justify-center rounded-2xl border bg-[url("/images/block/circles.svg")] bg-cover bg-center px-8 py-20 text-center md:p-20'>
            <div className="mx-auto max-w-screen-md">
              <h1 className="mb-4 text-balance text-3xl font-semibold md:text-5xl">
                Email enviado com sucesso
              </h1>
              <p className="text-muted-foreground md:text-lg">
                Enviamos um email com instruções, por favor verifique sua caixa de entrada.
              </p>
              <div className="mt-11 flex flex-col justify-center gap-2 sm:flex-row">
                <Button 
                  onClick={() => navigate("/")}
                  size="lg"
                >
                  Ir para a página inicial
                </Button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}