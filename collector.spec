%{!?BUILD_NUMBER: %define BUILD_NUMBER 1}
Name:           collector
Version:        1.0.0
Release:        080301
Group:          Applications/Internet
Summary:        Binary data collector
License:        MIT License
URL:            https://morphs.ru
Source0:        collector.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-root-%(%{__id_u} -n)
Requires:       daemonize
Requires(pre):  shadow-utils

# pull in golang libraries by explicit import path, inside the meta golang()
# [...]

%description
# include your full description of the application here.

%prep
%setup -q -n %{name}

# many golang binaries are "vendoring" (bundling) sources, so remove them. Those dependencies need to be packaged independently.
rm -rf vendor

%build
export GOPATH=$(pwd)
export 
go get -d
go build -v -a -ldflags "-B 0x$(head -c20 /dev/urandom|od -An -tx1|tr -d ' \n')" -tags 'netgo'

%install
rm -rf %{buildroot}
install -d %{buildroot}
install -d %{buildroot}%{_bindir}
install -d %{buildroot}%{_sysconfdir}
install -d  %{buildroot}%{_sysconfdir}/%{name}
install -d  %{buildroot}%{_sysconfdir}/init.d
install -d  %{buildroot}/var/log/%{name}
install -p -m 0755 ./%{name} %{buildroot}%{_bindir}/%{name}
install -p -m 0755 ./etc/%{name}.json %{buildroot}%{_sysconfdir}/%{name}/%{name}.json.example
install -p -m 0755 ./etc/init.d %{buildroot}%{_sysconfdir}/init.d/%{name}

%files
%defattr(-,root,root,-)
%attr(0755,root,root) %{_bindir}/%{name}
%attr(0755,root,root) %{_sysconfdir}/%{name}/%{name}.json.example
%attr(0755,root,root) %{_sysconfdir}/init.d/%{name}.example
# %{_sysconfdir}/init.d/%{name}

%pre

%changelog

* Thu Aug 03 2017 Mikhail Kirillov <mikkirillov@yandex.ru> - 1.0.0
 - first package version

