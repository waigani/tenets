// A Hello World! program in C#.
using System;
namespace HelloWorld
{
    class Hello 
    {
        // 
        static void Main() {
            Console.WriteLine(PlusEqualsOverridesParam("Starting String"));
        }

        // Positive Cases 

        static string PlusEqualsOverridesParam(string stringArg) {
            for (int i = 1; i <= 10; i++) {
                stringArg += i;
            }
            return stringArg;
        }

        static string PlusEqualsOverridesVar() {
            var stringArg = "const";
            for (int i = 1; i <= 10; i++) {
                stringArg += i;
            }
            return stringArg;
        }

        static string NestedOverrideParam(string stringArg) {
            stringArg = "someString";
            return stringArg;
        }

        // Negative cases
        static string FlatOverrideVar() {
            var stringArg = "someString";
            stringArg = "someOtherString";
            return stringArg;
        }

        static string FlatOverrideParam(string stringArg) {
            stringArg = "someString";
            return stringArg;
        }

        static string NonUsingOverriding(string stringArg) {
            stringArg = "overridden";
            return "Another string";
        }
        static string NonOverridingUsing(string stringArg) {
            return stringArg;
        }

        
    }
}
