ALTER TABLE "Payment" DROP CONSTRAINT "user_payment";
ALTER TABLE "Todo" DROP CONSTRAINT "user_todo";

DROP TABLE "User";
DROP TABLE "Todo";
DROP TABLE "Payment";
DROP TABLE "Quota";
DROP TABLE "PremiumType";

DROP TYPE payment_status;

