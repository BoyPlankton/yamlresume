# yamlresume: Build and Maintain Your Resume in YAML

**yamlresume** is a Golang-based tool that lets you write your resume in a simple YAML file and compile it into clean, professional HTML or Markdown output. This approach offers several benefits:

* **Version Control:**  Store your resume data in a YAML file, allowing you to track changes using Git and collaborate with GitOps workflows.
* **Flexibility:**  YAML offers a clear and concise format for structuring your resume content. 
* **Multiple Outputs:** Generate HTML or Markdown versions of your resume, suitable for different use cases.
* **Single Binary:** yamlresume is distributed as a single binary for easy execution.

## Install

## Create your resume YAML file
   Save your resume content in a file named `resume.yaml` (or any name you prefer).  Use the following structure as a guide, which outlines the supported schema for your resume data:

### Resume Schema

The YAML schema defines the structure for your resume content. Here's a breakdown of the available sections and their properties:

#### Basics
    * `name` (string): Your full name
    * `label` (string - optional): Your job title or professional label (e.g., Security Engineer)
    * `email` (string): Your email address
    * `phone` (string - optional): Your phone number in a preferred format
    * `url` (string - optional): Link to your personal website or online portfolio
    * `summary` (string, multi-line):  A brief and compelling summary of your experience and qualifications.
    * `location` (object - optional): Your physical address
        * `address` (string): Street address
        * `postalCode` (string): Postal code
        * `city` (string): City
        * `countryCode` (string - ISO 3166-1 alpha-2): Two-letter country code (e.g., US, CA)
        * `region` (string): State or province (e.g., CA for California)
    * `profiles` (list of objects - optional): Links to your online profiles
        * `network` (string): Social media platform name (e.g., Twitter, LinkedIn)
        * `username` (string): Username on the social media platform
        * `url` (string - optional): Link to your profile on the social media platform

#### Work Experience (listed in reverse chronological order)
    * `work` (list of objects):  An array containing entries for your work experience.
        * `name` (string):  Company name
        * `position` (string):  Your job title at the company
        * `url` (string - optional): Link to the company website
        * `startDate` (string):  Start date of employment in YYYY-MM format
        * `endDate` (string - optional): End date of employment in YYYY-MM format (if applicable)
        * `summary` (string, multi-line):  A brief description of your responsibilities and achievements at the company
        * `highlights` (list of strings):  Bullet points outlining key accomplishments and contributions

#### Volunteer Experience (follows the same structure as Work Experience)
    * `volunteer` (list of objects)

#### Education (listed in reverse chronological order)
    * `education` (list of objects):  An array containing entries for your educational background.
        * `institution` (string):  Name of the educational institution
        * `url` (string - optional): Link to the institution's website
        * `area` (string):  Field of study (e.g., Computer Science)
        * `studyType` (string):  Type of degree (e.g., Bachelor of Science, Master of Arts)
        * `graduationDate` (string): Graduation date in YYYY-MM format
        * `score` (string - optional):  Academic achievement score (e.g., GPA)

#### Awards & Recognition
    * `awards` (list of objects):  A list of awards and recognitions you have received
        * `title` (string):  Name of the award
        * `date` (string):  Date the award was received in YYYY-MM format
        * `awarder` (string):  Name of the institution or organization awarding the recognition
        * `summary` (string):  A brief description of the award

#### Certifications
    * `certificates` (list of objects):  A list of your professional certifications
        * `name` (string):  Name of the certification
        
### Example YAML file

```
basics:
  name: John Doe
  label: Security Engineer
  email: john.doe@example.com
  phone: (912) 555-4321
  url: https://johndoe.com
  summary: Highly motivated and results-oriented Computer Security Engineer with 12+ years of experience in safeguarding IT infrastructure. Proven ability to identify, analyze, and mitigate security risks, develop and implement security policies, and conduct vulnerability assessments and penetration testing. Strong understanding of security best practices and emerging cyber threats. Excellent communication and collaboration skills.
  location:
    address: 2712 Broadway St
    postalCode: 94155
    city: San Francisco
    countryCode: US
    region: CA
  profiles:
    - network: Twitter
      username: john
      url: https://twitter.com/john
work:
  - name: ABC Company
    position: Senior Security Engineer
    url: https://abccompany.com
    startDate: 2020-01-01
    summary: Description ...
    highlights:
      - Led vulnerability assessment and penetration testing (VAPT) engagements for critical infrastructure, identifying and remediating high-risk vulnerabilities.
      - Developed and implemented comprehensive security policies and procedures to enhance overall security posture.
      - Designed and deployed security controls such as firewalls, intrusion detection/prevention systems (IDS/IPS), and endpoint security solutions.
      - Conducted ongoing threat intelligence monitoring and analysis to stay abreast of emerging threats and implement preventative measures.
      - Managed and responded to security incidents, performing root cause analysis, containment, and eradication.
      - Collaborated with cross-functional teams (IT Operations, Development) to ensure security best practices are integrated into processes.
  - name: DEF Company
    position: Security Engineer
    url: https://defcompany.com
    startDate: 2015-01-01
    endDate: 2020-01-01
    summary: Description ...
    highlights:
      - Conducted vulnerability assessments and penetration testing for internal applications and systems.
      - Supported incident response activities, including investigation, containment, and recovery.
      - Monitored network activity for suspicious behavior and maintained security information and event management (SIEM) systems.
      - Provided security awareness training to employees on various cybersecurity topics.
      - Maintained and updated security documentation including security policies, procedures, and incident response plans.
volunteer:
  - organization: Organization
    position: Volunteer
    url: https://organization.com/
    startDate: 2012-01-01
    endDate: 2013-01-01
    summary: Description ...
    highlights:
      - Awarded volunteer of the month.
education:
  - institution: XYZ University
    url: https://xyzuniversity.edu/
    area: Computer Science (Security Concentration)
    studyType: Bachelor of Science
    graduationDate: 2015-01-01
    score: 4.0
awards:
  - title: Cybersecurity Excellence Award - Cloud Security
    date: 2023-11-01
    awarder: Cyber Defense Magazine
    summary: This award recognizes John Doe's outstanding contributions in securing cloud environments.
  - title: Runner-Up - Best Internal Penetration Test Report
    date: 2022-11-01
    awarder: ABC Company Internal Security Committee
    summary: This award highlights John Doe's exceptional skills in identifying and documenting vulnerabilities within an internal network.
certificates:
  - name: Certified Ethical Cloud Security Specialist (CECSS)
    issueDate: 2022-08-01
    expireDate: 2025-08-01
    certificateID: CSPA-CECSS-12345)
    issuer: Cyber Security Professionals Association (CSPA)
    url: https://certificate.com
  - name: Advanced Threat Detection and Analysis (ATDA)
    issueDate: 2021-05-01
    expireDate: 2024-05-01
    certificateID: ICR-ATDA-54321)
    issuer: Institute for Cybercrime Research (ICR)
    url: https://certificate.com
publications:
  - name: Publication
    publisher: Company
    releaseDate: 2014-10-01
    url: https://publication.com
    summary: Description ...
skills:
  - name: Network Security
    level: Master
    keywords:
      - Firewalls
      - IDS/IPS
  - name: Security Information and Event Management
    level: Master
    keywords:
      - Splunk
  - name: Cloud Security
    level: Master
    keywords:
      - AWS
      - Azure
      - GCP
languages:
  - language: English
    fluency: Native Speaker
interests:
  - name: Wildlife
    keywords:
      - Ferrets
      - Unicorns
references:
  - name: Jane Doe
    reference: Reference ...
projects:
  - name: Project
    startDate: 2019-01-01
    endDate: 2021-01-01
    description: Description ...
    highlights:
      - Won award at AIHacks 2016
    url: https://project.com
```