
	package main
	//
	func Create(){
	query := "INSERT INTO payment(id,order_id,transaction_id,payment_type,payment_method,payment_status,amount,admin_fee,device_i_d,msisdn_sender,msisdn_receiver,product_name,billing_number,purchase_mode)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err, r.log

	}
	_, err = stmt.Exec(payment.Id,payment.OrderId,payment.TransactionId,payment.PaymentType,payment.PaymentMethod,payment.PaymentStatus,payment.Amount,payment.AdminFee,payment.DeviceID,payment.MsisdnSender,payment.MsisdnReceiver,payment.ProductName,payment.BillingNumber,payment.PurchaseMode)
	if err != nil {
		r.log.Message += "|Exec|" + err.Error()
		return err, r.log
	}
	
	}