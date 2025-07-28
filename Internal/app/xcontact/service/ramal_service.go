package service

import (
	"ProjetoGustavo/Internal/app/xcontact/dto"
	"ProjetoGustavo/Internal/app/xcontact/model"
	"ProjetoGustavo/Internal/app/xcontact/repositories"
)

var AdicionarRamal = func(request dto.RamalRequest) (dto.RamalResponse, error) {
	ramal := model.Ramal{
		Numero:           request.Numero,
		Nome:             request.Nome,
		Senha:            request.Senha,
		Grupo:            request.Grupo,
		Allow:            request.Allow,
		Insecure:         request.Insecure,
		SubscribeContext: request.SubscribeContext,
		PickupGroup:      request.PickupGroup,
		CallGroup:        request.CallGroup,
		Transport:        request.Transport,
		CallLimit:        request.CallLimit,
		Nat:              request.Nat,
		Mac:              request.Mac,
		AccountCode:      request.AccountCode,
		DtmfMode:         request.DtmfMode,
		Language:         request.Language,
		MusicOnHold:      request.MusicOnHold,
		CallerID:         request.CallerID,
		HasSIP:           request.HasSIP,
		Encryption:       request.Encryption,
		Avpf:             request.Avpf,
		IceSupport:       request.IceSupport,
		DtlsEnable:       request.DtlsEnable,
		DtlsVerify:       request.DtlsVerify,
		DtlsCertFile:     request.DtlsCertFile,
		DtlsSetup:        request.DtlsSetup,
		RtcpMux:          request.RtcpMux,
		XcTipo:           request.XcTipo,
		TemMulti:         request.TemMulti,
		MultiRegra:       request.MultiRegra,
	}
	//RamalMulti:       request.RamalMulti,

	err := repositories.AdicionarRamal(ramal)

	if err != nil {
		return dto.RamalResponse{}, err
	}

	return dto.RamalResponse{
		Id:               ramal.Id,
		Numero:           ramal.Numero,
		Nome:             ramal.Nome,
		Senha:            ramal.Senha,
		Grupo:            ramal.Grupo,
		HasSIP:           ramal.HasSIP,
		Allow:            ramal.Allow,
		Insecure:         &ramal.Insecure,
		SubscribeContext: ramal.SubscribeContext,
		PickupGroup:      ramal.PickupGroup,
		CallGroup:        ramal.CallGroup,
		Transport:        ramal.Transport,
		CallLimit:        ramal.CallLimit,
		Nat:              ramal.Nat,
		Mac:              &ramal.Mac,
		AccountCode:      &ramal.AccountCode,
		DtmfMode:         &ramal.DtmfMode,
		Language:         ramal.Language,
		MusicOnHold:      &ramal.MusicOnHold,
		CallerID:         &ramal.CallerID,
		Encryption:       ramal.Encryption,
		Avpf:             ramal.Avpf,
		IceSupport:       ramal.IceSupport,
		DtlsEnable:       ramal.DtlsEnable,
		DtlsVerify:       &ramal.DtlsVerify,
		DtlsCertFile:     &ramal.DtlsCertFile,
		DtlsSetup:        &ramal.DtlsSetup,
		RtcpMux:          ramal.RtcpMux,
		XcTipo:           ramal.XcTipo,
		TemMulti:         ramal.TemMulti,
		MultiRegra:       &ramal.MultiRegra,
	}, nil

}

var ListarRamais = func() ([]dto.RamalResponse, error) {
	ramais, erro := repositories.ListarRamais()

	if erro != nil {
		return nil, erro
	}

	var response []dto.RamalResponse

	for _, ramal := range ramais {
		response = append(response, dto.RamalResponse{
			Id:               ramal.Id,
			Numero:           ramal.Numero,
			Nome:             ramal.Nome,
			Senha:            ramal.Senha,
			Grupo:            ramal.Grupo,
			HasSIP:           ramal.HasSIP,
			Allow:            ramal.Allow,
			Insecure:         &ramal.Insecure,
			SubscribeContext: ramal.SubscribeContext,
			PickupGroup:      ramal.PickupGroup,
			CallGroup:        ramal.CallGroup,
			Transport:        ramal.Transport,
			CallLimit:        ramal.CallLimit,
			Nat:              ramal.Nat,
			Mac:              &ramal.Mac,
			AccountCode:      &ramal.AccountCode,
			DtmfMode:         &ramal.DtmfMode,
			Language:         ramal.Language,
			MusicOnHold:      &ramal.MusicOnHold,
			CallerID:         &ramal.CallerID,
			Encryption:       ramal.Encryption,
			Avpf:             ramal.Avpf,
			IceSupport:       ramal.IceSupport,
			DtlsEnable:       ramal.DtlsEnable,
			DtlsVerify:       &ramal.DtlsVerify,
			DtlsCertFile:     &ramal.DtlsCertFile,
			DtlsSetup:        &ramal.DtlsSetup,
			RtcpMux:          ramal.RtcpMux,
			XcTipo:           ramal.XcTipo,
			TemMulti:         ramal.TemMulti,
			MultiRegra:       &ramal.MultiRegra})

	}

	return response, nil
}

var BuscarRamalPorId = func(id int) (dto.RamalResponse, error) {

	ramal, erro := repositories.BuscarRamalPorId(id)

	if erro != nil {
		return dto.RamalResponse{}, erro
	}

	return dto.RamalResponse{
		Id:               ramal.Id,
		Numero:           ramal.Numero,
		Nome:             ramal.Nome,
		Senha:            ramal.Senha,
		Grupo:            ramal.Grupo,
		HasSIP:           ramal.HasSIP,
		Allow:            ramal.Allow,
		Insecure:         &ramal.Insecure,
		SubscribeContext: ramal.SubscribeContext,
		PickupGroup:      ramal.PickupGroup,
		CallGroup:        ramal.CallGroup,
		Transport:        ramal.Transport,
		CallLimit:        ramal.CallLimit,
		Nat:              ramal.Nat,
		Mac:              &ramal.Mac,
		AccountCode:      &ramal.AccountCode,
		DtmfMode:         &ramal.DtmfMode,
		Language:         ramal.Language,
		MusicOnHold:      &ramal.MusicOnHold,
		CallerID:         &ramal.CallerID,
		Encryption:       ramal.Encryption,
		Avpf:             ramal.Avpf,
		IceSupport:       ramal.IceSupport,
		DtlsEnable:       ramal.DtlsEnable,
		DtlsVerify:       &ramal.DtlsVerify,
		DtlsCertFile:     &ramal.DtlsCertFile,
		DtlsSetup:        &ramal.DtlsSetup,
		RtcpMux:          ramal.RtcpMux,
		XcTipo:           ramal.XcTipo,
		TemMulti:         ramal.TemMulti,
		MultiRegra:       &ramal.MultiRegra,
	}, nil
}

var AtualizarRamal = func(id int, request dto.RamalRequest) (dto.RamalResponse, error) {
	ramal := model.Ramal{
		Numero:           request.Numero,
		Nome:             request.Nome,
		Senha:            request.Senha,
		Grupo:            request.Grupo,
		Allow:            request.Allow,
		Insecure:         request.Insecure,
		SubscribeContext: request.SubscribeContext,
		PickupGroup:      request.PickupGroup,
		CallGroup:        request.CallGroup,
		Transport:        request.Transport,
		CallLimit:        request.CallLimit,
		Nat:              request.Nat,
		Mac:              request.Mac,
		AccountCode:      request.AccountCode,
		DtmfMode:         request.DtmfMode,
		Language:         request.Language,
		MusicOnHold:      request.MusicOnHold,
		CallerID:         request.CallerID,
		HasSIP:           request.HasSIP,
		Encryption:       request.Encryption,
		Avpf:             request.Avpf,
		IceSupport:       request.IceSupport,
		DtlsEnable:       request.DtlsEnable,
		DtlsVerify:       request.DtlsVerify,
		DtlsCertFile:     request.DtlsCertFile,
		DtlsSetup:        request.DtlsSetup,
		RtcpMux:          request.RtcpMux,
		XcTipo:           request.XcTipo,
		TemMulti:         request.TemMulti,
		MultiRegra:       request.MultiRegra,
	}
	//RamalMulti:       request.RamalMulti,

	erro := repositories.AtualizarRamal(id, ramal)

	if erro != nil {
		return dto.RamalResponse{}, erro
	}

	return dto.RamalResponse{
		Id:               ramal.Id,
		Numero:           ramal.Numero,
		Nome:             ramal.Nome,
		Senha:            ramal.Senha,
		Grupo:            ramal.Grupo,
		HasSIP:           ramal.HasSIP,
		Allow:            ramal.Allow,
		Insecure:         &ramal.Insecure,
		SubscribeContext: ramal.SubscribeContext,
		PickupGroup:      ramal.PickupGroup,
		CallGroup:        ramal.CallGroup,
		Transport:        ramal.Transport,
		CallLimit:        ramal.CallLimit,
		Nat:              ramal.Nat,
		Mac:              &ramal.Mac,
		AccountCode:      &ramal.AccountCode,
		DtmfMode:         &ramal.DtmfMode,
		Language:         ramal.Language,
		MusicOnHold:      &ramal.MusicOnHold,
		CallerID:         &ramal.CallerID,
		Encryption:       ramal.Encryption,
		Avpf:             ramal.Avpf,
		IceSupport:       ramal.IceSupport,
		DtlsEnable:       ramal.DtlsEnable,
		DtlsVerify:       &ramal.DtlsVerify,
		DtlsCertFile:     &ramal.DtlsCertFile,
		DtlsSetup:        &ramal.DtlsSetup,
		RtcpMux:          ramal.RtcpMux,
		XcTipo:           ramal.XcTipo,
		TemMulti:         ramal.TemMulti,
		MultiRegra:       &ramal.MultiRegra,
	}, nil
}

var DeletarRamal = func(id int) error {
	erro := repositories.DeletarRamal(id)

	if erro != nil {
		return erro
	}
	return nil
}
