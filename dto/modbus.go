package dto

type CommissioningToolRequest struct {
	NetworkUUID     string `json:"network_uuid"`
	DeviceUUID      string `json:"device_uuid"`
	RegisterType    string `json:"register_type"`
	StartingAddress uint16 `json:"starting_address"`
	RegisterCount   uint16 `json:"register_count"`
}

type CommissioningToolResponse struct {
	Error        string                         `json:"error"`
	PollRequest  CommissioningToolPollRequest   `json:"poll_request"`
	PollResponse *CommissioningToolPollResponse `json:"poll_response"`
}

type CommissioningToolPollRequest struct {
	Type                          string `json:"type"`
	TCPHostIP                     string `json:"tcp_host_ip"`
	TCPPort                       string `json:"tcp_port"`
	SerialPort                    string `json:"serial_port"`
	SerialBaudRate                string `json:"serial_baud_rate"`
	SerialParity                  string `json:"serial_parity"`
	SerialStopBit                 string `json:"serial_stop_bit"`
	DeviceAddress                 string `json:"device_address"`
	ModbusRegisterType            string `json:"modbus_register_type"`
	ModbusRegisterStartingAddress uint16 `json:"modbus_register_starting_address"`
	ModbusRegisterCount           uint16 `json:"modbus_register_count"`
}

type CommissioningToolPollResponse struct {
	RawBytes            *CommissioningToolRawByte              `json:"raw_bytes"`
	DataInterpretations []*CommissioningToolDataInterpretation `json:"data_interpretations"`
}

type CommissioningToolRawByte struct {
	Hex     []string  `json:"hex"`
	Decimal []float64 `json:"decimal"`
}

type CommissioningToolDataInterpretation struct {
	Endianness      string                             `json:"endianness"`
	EndiannessName  string                             `json:"endianness_name"`
	Interpretations []*CommissioningToolInterpretation `json:"interpretations"`
}

type CommissioningToolInterpretation struct {
	DateType string      `json:"date_type"`
	Values   interface{} `json:"values"`
}
