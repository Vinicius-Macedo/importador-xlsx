package services

import (
	"api/cmd/internal/postgresrepo"
	"context"
	"log"
)

func (s *Service) CreatePartnersWithCopy(params []postgresrepo.CreatePartnerParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreatePartner(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateCustomersWithCopy(params []postgresrepo.CreateCustomerParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateCustomer(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateSkusWithCopy(params []postgresrepo.CreateSkuParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateSku(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateProductsWithCopy(params []postgresrepo.CreateProductParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateProduct(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreatePublishersWithCopy(params []postgresrepo.CreatePublisherParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreatePublisher(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateSubscriptionsWithCopy(params []postgresrepo.CreateSubscriptionParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateSubscription(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateMetersWithCopy(params []postgresrepo.CreateMeterParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateMeter(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateResourcesWithCopy(params []postgresrepo.CreateResourceParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateResource(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateEntitlementsWithCopy(params []postgresrepo.CreateEntitlementParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateEntitlement(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateBenefitsWithCopy(params []postgresrepo.CreateBenefitParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateBenefit(ctx, params)

	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *Service) CreateBillingsWithCopy(params []postgresrepo.CreateBillingParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultDBTimeout)
	defer cancel()

	value, err := s.queries.CreateBilling(ctx, params)

	if err != nil {
		log.Println("Error inserting billing record", err)
		return 0, err
	}

	return value, nil
}
