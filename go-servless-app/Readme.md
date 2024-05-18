zip the executable command
zip function.zip main

aws lambda create-function --function-name go-servless-app \
--zip-file fileb://function.zip --handler main --runtime go1.x \
--role arn:aws:iam::303580349437:role/lambda-ex



aws lambda invoke --function-name go-servless-app --cli-binary-format raw-in-base64-out --payload '{"What is your name ?" :"Jim","How old are you?":20}' op.txt


Benefits 

1)Servless 
2)Seamlessly scale bases upon traffic 
3)Lambda response to event stream
