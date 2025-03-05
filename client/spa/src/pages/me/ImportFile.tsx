import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Container } from "@/components/ui/container";
import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";

interface ImportFileProps {}

export function ImportFile(props: ImportFileProps) {
  const [file, setFile] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  const mutation = useMutation({
    mutationFn: async () => {
      const response = await axiosInstance.post(
        "/import",
        {
          file: file,
        },
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      return response.data;
    },
    onSuccess: () => {
      setFile(null);
      setSuccess(true);
      setError(null);
    },
    onError: () => {
      setError("Erro ao enviar o arquivo, contate o suporte");
    },
  });

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setFile(event.target.files[0]);
    }
  };

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    if (file) {
      mutation.mutate;
    } else {
      setError("Por favor, selecione um arquivo");
    }
  };

  return (
    <Container>
      <div className="bg-white p-8 rounded-lg shadow-md w-full">
        <h1 className="text-2xl font-bold mb-4">Importar Arquivo xlsx</h1>
        {error && (
          <div
            className="bg-red-100 border border-red-400 text-red-700 px-4 py-2 rounded relative mb-4"
            role="alert"
          >
            <span className="block sm:inline">{error}</span>
          </div>
        )}
        {success && (
          <div
            className="bg-green-100 border border-green-400 text-green-700 px-4 py-2 rounded relative mb-4"
            role="alert"
          >
            <span className="block sm:inline">Arquivo enviado com sucesso</span>
          </div>
        )}
        <p className="mb-4">
          Por favor, selecione um arquivo xlsx com as seguintes colunas:
        </p>
        <ul className="list-disc list-inside mb-4 grid lg:grid-cols-3">
          <li>PartnerId</li>
          <li>PartnerName</li>
          <li>CustomerId</li>
          <li>CustomerName</li>
          <li>CustomerDomainName</li>
          <li>CustomerCountry</li>
          <li>MpnId</li>
          <li>Tier2MpnId</li>
          <li>InvoiceNumber</li>
          <li>ProductId</li>
          <li>SkuId</li>
          <li>AvailabilityId</li>
          <li>SkuName</li>
          <li>ProductName</li>
          <li>PublisherName</li>
          <li>PublisherId</li>
          <li>SubscriptionDescription</li>
          <li>SubscriptionId</li>
          <li>ChargeStartDate</li>
          <li>ChargeEndDate</li>
          <li>UsageDate</li>
          <li>MeterType</li>
          <li>MeterCategory</li>
          <li>MeterId</li>
          <li>MeterSubCategory</li>
          <li>MeterName</li>
          <li>MeterRegion</li>
          <li>Unit</li>
          <li>ResourceLocation</li>
          <li>ConsumedService</li>
          <li>ResourceGroup</li>
          <li>ResourceURI</li>
          <li>ChargeType</li>
          <li>UnitPrice</li>
          <li>Quantity</li>
          <li>UnitType</li>
          <li>BillingPreTaxTotal</li>
          <li>BillingCurrency</li>
          <li>PricingPreTaxTotal</li>
          <li>PricingCurrency</li>
          <li>ServiceInfo1</li>
          <li>ServiceInfo2</li>
          <li>Tags</li>
          <li>AdditionalInfo</li>
          <li>EffectiveUnitPrice</li>
          <li>PCToBCExchangeRate</li>
          <li>PCToBCExchangeRateDate</li>
          <li>EntitlementId</li>
          <li>EntitlementDescription</li>
          <li>PartnerEarnedCreditPercentage</li>
          <li>CreditPercentage</li>
          <li>CreditType</li>
          <li>BenefitOrderId</li>
          <li>BenefitId</li>
          <li>BenefitType</li>
        </ul>
        <form
          onSubmit={handleSubmit}
          className="space-y-4 flex flex-col items-start"
        >
          <div>
            <Label htmlFor="file">Selecione o arquivo xlsx</Label>
            <Input
              id="file"
              type="file"
              accept=".xlsx"
              onChange={handleFileChange}
              className="mt-2"
            />
          </div>
          <Button
            type="submit"
            className="w-full lg:max-w-[295px]"
            disabled={mutation.isPending}
            onClick={mutation.mutate}
          >
            {mutation.isPending ? <div className="spinner"></div> : "Enviar"}
          </Button>
        </form>
      </div>
    </Container>
  );
}
