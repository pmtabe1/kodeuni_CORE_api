import java.util.Scanner;
import java.lang.Math;


public class cashRegister {


	public static void main(String[] args) {
	// get input from cashier	
	Scanner input = new Scanner(System.in);
	
	// initialize variables
	double total = 0.0, coupontotal = 0.0; // initialize totals
	double price, coupon, cash; // initialize input variables
	double subtotal, taxpercentage, totalwithtax, finaltotal, cashowed, amount, cashleft; // initialize calculation variables
	int itemcounter = 0, couponcounter = 0; // initialize counters
	int pennies, nickels, dimes, quarters, ones, fives, tens, twenties;
	
	// CONSTANT VARIABLES
		double TAXAMOUNT = .06;
		double PENNIES = .01;
		double NICKELS = .05;
		double DIMES = .1;
		double QUARTERS = .25;
		double DOLLARS = 1;
		double FIVES = 5;
		double TENS = 10;
		double TWENTIES = 20;
		
		
	// processing phase
	// prompt cashier for input and read price from user
	System.out.print("Enter price or 0 to quit: ");
	price = input.nextDouble();
	
	// loop until sentinel value read from user
	while (price != 0)
	{
		total = total + price; // add item price to total
		itemcounter = itemcounter + 1;
		
		// prompt for input and read next price from cashier
		System.out.print("Enter price or 0 to quit: ");
		price = input.nextDouble();
	}
	
	
	
	// termination phase
	// if cashier entered at least one item price
	if (itemcounter != 0)
	{
		// prompt user to ask for coupons
		System.out.print("Please enter any coupon amounts at this time, or press 0 to continue: ");
		coupon = input.nextDouble();
		
		// loop until sentinel value read from user
		while (coupon != 0)
		{
			coupontotal = coupontotal + coupon; // add coupon total
			couponcounter = couponcounter + 1;
			
			// prompt for input and read next coupon from cashier
			System.out.print("Enter coupon amount or 0 to quit: ");
			coupon = input.nextDouble();
		}
		
		// Display item total, coupon total, subtotal, tax percentage, tax amount and total
		
		System.out.printf("ITEM TOTAL: %.2f%n", total);
		System.out.printf("\n" + "COUPON TOTAL: %.2f%n", coupontotal);
		
		subtotal = total - coupontotal;
		System.out.printf("\n" + "SUBTOTAL: %.2f%n", subtotal);
		
		taxpercentage = TAXAMOUNT * 100;
		System.out.printf("\n" + "TAX PERCENTAGE: %.2f%n", taxpercentage);
		
		totalwithtax = TAXAMOUNT * subtotal;
		System.out.printf("\n" + "TAX AMOUNT: %.2f%n", totalwithtax);
		
		finaltotal = totalwithtax + subtotal;
		System.out.printf("\n" + "TOTAL: %.2f%n", finaltotal);
		
		// Prompt cashier to enter cash given by customer
		System.out.print("\n" + "Enter amount of cash given: ");
		cash = input.nextDouble();
		
		// Tell cashier the amount to give back to customer
		cashowed = cash - finaltotal;
		System.out.printf("\n" + "CHANGE OWED: %.2f%n", cashowed);
		
		// Display amount of each bill to give back in change
     
		twenties = (int) Math.floor(cashowed/TWENTIES);
		tens = (int) Math.floor((cashowed - twenties * TWENTIES)/TENS);
		fives = (int) Math.floor((cashowed-twenties * TWENTIES - tens * TENS)/FIVES);
		ones = (int) Math.floor((cashowed - twenties * TWENTIES - tens * TENS - fives * FIVES));
        
		cashleft = cashowed - twenties * TWENTIES - tens * TENS - fives * FIVES - ones * DOLLARS;
		
		// correct incorrect decimal of .999999
		cashleft = Math.round(cashleft * 100.0) /100.0;
		
		quarters =  (int) Math.floor(cashleft/.25);
		
        cashleft = cashowed - twenties * TWENTIES - tens * TENS - fives * FIVES - ones * DOLLARS - quarters * QUARTERS;
		
		// correct incorrect decimal of .999999
		cashleft = Math.round(cashleft * 100.0) /100.0;
		
		dimes =  (int) Math.floor(cashleft/DIMES);
		
		cashleft = cashowed - twenties * TWENTIES - tens * TENS - fives * FIVES - ones * DOLLARS - quarters * QUARTERS - dimes * DIMES;
		
		// correct incorrect decimal of .999999
		cashleft = Math.round(cashleft * 100.0) /100.0;
		
		nickels =  (int) Math.floor(cashleft/NICKELS);
		cashleft = cashowed - twenties * TWENTIES - tens * TENS - fives * FIVES - ones * DOLLARS - quarters * QUARTERS - dimes * DIMES - nickels * NICKELS;
		
		// correct incorrect decimal of .999999
		cashleft = Math.round(cashleft * 100.0) /100.0;
		
		pennies =  (int) Math.floor(cashleft/PENNIES);
		
	// Print correct change to give back	
	if (twenties > 0)
	{
		System.out.print("$20.00: " + twenties + "\n");
	}
		
	if (tens > 0)
	{
		System.out.print("$10.00: " + tens + "\n");
	}
		
	if (fives > 0)
	{
		System.out.print("$5.00: " + fives + "\n");
	}	
	
	if (ones > 0)
	{
		System.out.print("1.00: " + ones + "\n");
	}	
	
	if (quarters > 0)
	{
		System.out.print("$0.25: " + quarters + "\n");
	}	
		
	if (dimes > 0)
	{
		System.out.print("$0.10: " + dimes + "\n");
	}
		
	if (nickels > 0)
	{
		System.out.print("$0.05: " + nickels + "\n");
	}
		
	if (pennies > 0)
	{
		System.out.print("$0.01: " + pennies + "\n");
	}
	
		System.out.println("----------------------");
		System.out.println("***END OF PROGRAM***");
		
	
	}
	
	}			
 
}	