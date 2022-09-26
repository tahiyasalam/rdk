package dimensionengineering

type (
	commandCode byte
	opCode      byte
)

const (
	minSpeed = 0x0
	maxSpeed = 0x7f

	// Commands.
	singleForward   commandCode = 0
	singleBackwards commandCode = 1
	singleDrive     commandCode = 2
	multiForward    commandCode = 3
	multiBackward   commandCode = 4
	multiDrive      commandCode = 5
	multiTurnLeft   commandCode = 6
	multiTurnRight  commandCode = 7
	multiTurn       commandCode = 8

	// Serial level op-codes.
	opMotor1Forward       opCode = 0x00
	opMotor1Backwards     opCode = 0x01
	opMinVoltage          opCode = 0x02
	opMaxVoltage          opCode = 0x03
	opMotor2Forward       opCode = 0x04
	opMotor2Backwards     opCode = 0x05
	opMotor1Drive         opCode = 0x06
	opMotor2Drive         opCode = 0x07
	opMultiDriveForward   opCode = 0x08
	opMultiDriveBackwards opCode = 0x09
	opMultiDriveRight     opCode = 0x10
	opMultiDriveLeft      opCode = 0x11
	opMultiDrive          opCode = 0x12
	opMultiTurn           opCode = 0x13
	opSerialTimeout       opCode = 0x14
	opSerialBaudRate      opCode = 0x15
	opRamping             opCode = 0x16
	opDeadband            opCode = 0x17
)
