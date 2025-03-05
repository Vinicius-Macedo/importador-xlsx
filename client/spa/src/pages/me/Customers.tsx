import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

interface CustomersProps {}

export function Customers(props: CustomersProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const { isLoading, isError, data } = useQuery({
    queryKey: ["customers"],
    queryFn: async () => {
      const response = await axiosInstance.get("/customers");
      return response.data; // Return the data directly
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error loading data</div>;
  }

  const filteredData = data ? data.filter((customer: any) =>
    customer.name.toLowerCase().includes(searchTerm.toLowerCase())
  ) : [];

  return (
    <div className="p-6 w-full responsive-table">
      <div className="mb-4">
        <Label htmlFor="search">Pesquisar Cliente</Label>
        <Input
          id="search"
          type="text"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          placeholder="Digite para pesquisar..."
          className="mt-2"
        />
      </div>
      <div className="overflow-x-auto">
        {filteredData.length > 0 ? (
          <Table className="w-full">
            <TableCaption>Lista de Clientes</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead>Nome</TableHead>
                <TableHead>Domínio</TableHead>
                <TableHead>País</TableHead>
                <TableHead className="text-right">Tier to MPN ID</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredData.map((cliente: any) => (
                <TableRow key={cliente.customer_key}>
                  <TableCell>{cliente.name}</TableCell>
                  <TableCell>{cliente.domain_name}</TableCell>
                  <TableCell>{cliente.country}</TableCell>
                  <TableCell className="text-right">
                    {cliente.tier_to_mpn_id}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TableCell colSpan={3}>Total</TableCell>
                <TableCell className="text-right">
                  {filteredData.length}
                </TableCell>
              </TableRow>
            </TableFooter>
          </Table>
        ) : (
          <div className="text-center py-4">Nenhum cliente encontrado.</div>
        )}
      </div>
    </div>
  );
}