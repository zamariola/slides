class Analyst {
    protected String name;
    protected String email;

    String getDivision() {
        return "dev-team";
    }

    String getContact() {
        return this.name + " - " + this.email + " - " + this.getDivision();
    }
}

class Manager extends Analyst {
    private String division;

    @Override
    String getDivision() {
        return this.division;
    }

    Manager(String division, String name, String email) {
        this.name = name;
        this.email = email;
        this.division = division;
    }
}

class MainClass {
    public static void main(String... args) {

        Manager m = new Manager("management", "Hugo", "h@g.com");
        System.out.println(m.getContact());

    }
}