import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export function Home() {
  const navigate = useNavigate();

  return (
    <>
      <div className="flex flex-col items-center justify-center max-w-7xl mx-auto p-8">
        <div className="">
          <h1 className="text-2xl font-bold mb-4">
            Teste Técnico para Desenvolvedor Golang Pleno - Importador de Dados
          </h1>
          <p className="mb-4">
            Bem-vindo a minha resposta do teste técnico! O objetivo deste
            desafio é avaliar habilidades em desenvolvimento backend e frontend,
            bem como a capacidade de criar um sistema eficiente e
            performático.
          </p>
          <p className="mb-4">
            A documentação da API pode ser acessada{" "}
            <a
              href="/api/swagger/index.html"
              className="text-blue-500 hover:underline"
              target="_blank"
            >
              clicando aqui
            </a> se estiver autenticado.
          </p>
          <h2 className="text-xl font-semibold mb-2">Descrição do Desafio</h2>
          <p className="mb-4">
            Você deverá desenvolver um importador de dados para uma base de
            dados PostgreSQL, utilizando a linguagem Go (Golang). Esse
            importador será responsável por processar e armazenar os dados de um
            arquivo XLSX (com 80.000 linhas e 55 colunas) fornecido. Durante a
            avaliação, serão observados critérios como normalização dos dados e
            performance do importador.
          </p>
          <p className="mb-4">
            Além disso, será necessário desenvolver uma API em Go, que deverá
            conter:
          </p>
          <ul className="list-disc list-inside mb-4">
            <li>Endpoint de autenticação</li>
            <li>
              Endpoints de consulta para acessar os dados importados no banco de
              dados PostgreSQL
            </li>
          </ul>
          <p className="mb-4">
            Como um diferencial, será valorizado o desenvolvimento de um
            front-end em React, que deverá apresentar:
          </p>
          <ul className="list-disc list-inside mb-4">
            <li>Indicadores totalizadores</li>
            <li>
              Agrupamentos de categorias, recursos, clientes e meses de cobrança
            </li>
          </ul>
          <h2 className="text-xl font-semibold mb-2">Entrega do Projeto</h2>
          <p className="mb-4">
            Para avaliação, a aplicação deverá ser publicada em um ambiente
            acessível via link, juntamente com a documentação de execução do
            projeto, contendo:
          </p>
          <ul className="list-disc list-inside mb-4">
            <li>Passos para rodar a aplicação</li>
            <li>Configuração do ambiente</li>
            <li>Endpoints disponíveis na API</li>
            <li>
              Qualquer outra informação relevante para entendimento e uso do
              sistema
            </li>
          </ul>
          <h2 className="text-xl font-semibold mb-2">Prazo</h2>
          <p className="mb-4">
            A entrega do projeto deve ser realizada em até uma semana.
          </p>
          <div className="flex gap-2 justify-center items-center pt-8">
            <Button onClick={() => navigate("/login")}>Entrar</Button>
            <Button onClick={() => navigate("/register")}>Registrar</Button>
          </div>
        </div>
      </div>
    </>
  );
}
