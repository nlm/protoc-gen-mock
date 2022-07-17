package main

// func ZipGen() error {
// 	var msg proto.Message = &demopb.CreatePersonRequest{}
// 	msg = pbutils.Populate(msg)
// 	data, err := protojson.MarshalOptions{
// 		UseEnumNumbers:  true,
// 		EmitUnpopulated: true,
// 		UseProtoNames:   true,
// 	}.Marshal(msg)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(string(data))
// 	var buf bytes.Buffer
// 	var gw = gzip.NewWriter(&buf)
// 	if _, err := gw.Write(data); err != nil {
// 		return err
// 	}
// 	if err := gw.Close(); err != nil {
// 		return err
// 	}
// 	fmt.Print("var a = []byte{")
// 	for i, b := range buf.Bytes() {
// 		if i%12 == 0 {
// 			fmt.Print("\n\t")
// 		}
// 		fmt.Printf("0x%02x, ", b)
// 	}
// 	fmt.Print("\n}\n")
// 	return nil
// }
